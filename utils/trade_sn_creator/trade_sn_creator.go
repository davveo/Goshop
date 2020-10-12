package trade_sn_creator

import (
	"Goshop/global/consts"
	"Goshop/model"
	"Goshop/utils/redis"
	"Goshop/utils/time_utils"
	"Goshop/utils/yml_config"
	"fmt"
	"log"
	"sync"
	"time"
)

const LOCK_NAME = "trade_sn_error_lock"

var (
	rds  = redis.GetRedisClient()
	lock sync.RWMutex
)

type TradeSnCreator struct {
}

func (tsc *TradeSnCreator) GenerateTradeSn() string {

	key := fmt.Sprintf("{%s}_", consts.TRADE_SN_CACHE_PREFIX)

	return tsc.generateSn(key)
}

func (tsc *TradeSnCreator) GenerateOrderSn() string {

	key := fmt.Sprintf("{%s}_", consts.ORDER_SN_CACHE_PREFIX)

	return tsc.generateSn(key)
}

func (tsc *TradeSnCreator) GeneratePayLogSn() string {

	key := fmt.Sprintf("{%s}_", consts.PAY_LOG_SN_CACHE_PREFIX)

	return tsc.generateSn(key)
}

func (tsc *TradeSnCreator) GenerateSmallChangeLogSn() string {

	key := fmt.Sprintf("{%s}_", consts.SMALL_CHANGE_CACHE_PREFIX)

	return tsc.generateSn(key)
}

func (tsc *TradeSnCreator) GenerateAfterSaleServiceSn() string {

	key := fmt.Sprintf("{%s}_", consts.AFTER_SALE_SERVICE_PREFIX)

	return tsc.generateSn(key)
}

func (tsc *TradeSnCreator) cleanCache() {
	timeStr := tsc.getYesterday()
	rds.Remove(fmt.Sprintf("%s_%s", consts.TRADE_SN_CACHE_PREFIX, timeStr))
	rds.Remove(fmt.Sprintf("%s_%s", consts.ORDER_SN_CACHE_PREFIX, timeStr))
}

func (tsc *TradeSnCreator) getYesterday() string {
	return time.Now().Format("20060102")
}

func (tsc *TradeSnCreator) generateSn(key string) string {
	// 通过Redis的自增来控制编号的自增
	// key 区分类型的主key，日期会连在这个key后面
	var sn string

	timeStr := tsc.getYesterday()
	redisKey := key + "_" + timeStr
	redisSignKey := key + "_" + timeStr + "_SIGN"

	//用当天的时间进行自增
	snCount := tsc.getSnCount(redisKey, redisSignKey)
	//预计每天订单不超过1百万单
	num := 1000000
	if snCount < num {
		sn = fmt.Sprintf("%s%d", "000000", snCount)
		sn = sn[len(sn)-6:]
	} else {
		sn = string(snCount)
	}
	sn = timeStr + sn
	return ""
}

func (tsc *TradeSnCreator) getSnCount(redisKey, redisSignKey string) int {
	var snCount = 0
	result := tsc.getRedisScriptResult(redisKey, redisSignKey)
	//如果为-1，说明缓存被击穿了
	if result == -1 {
		//上锁
		lock.RLock()
		//如果并发这里有等待取锁的操作，则有可能出现多次处理redis 击穿问题，所以要重复判断是否redis被击穿
		result = tsc.getRedisScriptResult(redisKey, redisSignKey)
		//如果为-1，说明缓存被击穿了
		if result == -1 {
			//从库中读取当天的订单数量
			snCount = tsc.countFromDB()
			snCount++
			//重置计数器
			redis.GetRedisClient().Put(redisKey, string(snCount), 0)
			result = tsc.getRedisScriptResult(redisKey, redisSignKey)
		}
		snCount = result
	}

	lock.RUnlock()

	return snCount
}

func (tsc *TradeSnCreator) getRedisScriptResult(redisKey, redisSignKey string) int {
	// TODO
	return -1
}

func (tsc *TradeSnCreator) countFromDB() (rows int) {
	var (
		sql = "select count(1) from es_order where create_time >= ? and create_time <= ? "
	)

	sqlType := yml_config.CreateYamlFactory().GetString("UseDbType")
	dbDriver := model.CreateBaseSqlFactory(sqlType)

	err := dbDriver.QueryRow(sql,
		time_utils.GetToDayOfStart(),
		time_utils.GetToDayOfEnd()).Scan(&rows)
	if err != nil {
		log.Println("sql.count 错误", err.Error())
	}
	return
}
