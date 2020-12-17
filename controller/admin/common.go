package admin

import (
	"Goshop/global/consts"
	"Goshop/model"
	"Goshop/utils/common"
	"Goshop/utils/store/ceph"
	"Goshop/utils/store/oss"
	"Goshop/utils/yml_config"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SiteShow(ctx *gin.Context) {
	config := model.CreateSettingFactory("").Get(consts.SITE)
	if config == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "获取数据失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, config)
}

func ListFocusPicture(ctx *gin.Context) {
	clientType := ctx.Query("client_type") //APP/WAP/PC

	data, err := model.CreateFocusPictureFactory("").List(clientType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func CreateFocusPicture(ctx *gin.Context) {

}

func UpdateFocusPicture(ctx *gin.Context) {

}

func FindOneFocusPicture(ctx *gin.Context) {

}

func DelFocusPicture(ctx *gin.Context) {

}

func FindOneClientPage(ctx *gin.Context) {
	pageType := ctx.Query("page_type")     // APP/WAP/PC
	clientType := ctx.Query("client_type") // INDEX/SPECIAL
	data, err := model.CreatePageFactory("").GetByType(clientType, pageType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func UpdateClientPage(ctx *gin.Context) {

}

func UpdatePage(ctx *gin.Context) {

}

func FindOnePage(ctx *gin.Context) {

}

func ListSiteNavigation(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	clientType := ctx.Query("client_type")
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["page_no"] = pageNo
	queryParams["page_size"] = pageSize
	queryParams["client_type"] = clientType
	data, dataTotal := model.CreateSiteNavigationFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func CreateSiteNavigation(ctx *gin.Context) {

}

func UpdateSiteNavigation(ctx *gin.Context) {

}

func DelSiteNavigation(ctx *gin.Context) {

}

func FindOneSiteNavigation(ctx *gin.Context) {

}

func UpdateSiteNavigationSort(ctx *gin.Context) {

}

func HotKeyWordsList(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["page_no"] = pageNo
	queryParams["page_size"] = pageSize
	data, dataTotal := model.CreateHotKeyWordFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func CreateHotKeyWords(ctx *gin.Context) {

}

func FindOneHotKeyWords(ctx *gin.Context) {

}

func UpdateHotKeyWords(ctx *gin.Context) {

}

func DelHotKeyWords(ctx *gin.Context) {

}

func AdminTask(ctx *gin.Context) {
	taskType := ctx.Param("task_type") // PAGE_CREATE/GOODS_INDEX_INIT
	fmt.Println(taskType)
}

func AdminTaskProgress(ctx *gin.Context) {

}

func DelAdminTask(ctx *gin.Context) {

}

func GoodsSearchCustomWord(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	keyWords := ctx.DefaultQuery("keywords", "")
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["page_no"] = pageNo
	queryParams["keywords"] = keyWords
	queryParams["page_size"] = pageSize

	data, dataTotal := model.CreateCustomWordsFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func CreateGoodsSearchCustomWord(ctx *gin.Context) {

}

func EditGoodsSearchCustomWord(ctx *gin.Context) {

}

func DelGoodsSearchCustomWord(ctx *gin.Context) {

}

func FindOneGoodsSearchCustomWord(ctx *gin.Context) {

}

func CreateEsCustomWordSecretKey(ctx *gin.Context) {

}

func FindEsCustomWordSecretKey(ctx *gin.Context) {

}

func GoodsSearchGoodsWord(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	keyWords := ctx.DefaultQuery("keywords", "")
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["page_no"] = pageNo
	queryParams["keywords"] = keyWords
	queryParams["page_size"] = pageSize

	data, dataTotal := model.CreateGoodsWordsFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}

func CreateGoodsSearchGoodsWord(ctx *gin.Context) {

}

func DelGoodsSearchGoodsWord(ctx *gin.Context) {

}

func EditGoodsSearchGoodsWord(ctx *gin.Context) {

}

func SortGoodsSearchGoodsWord(ctx *gin.Context) {

}

func GoodsSearchCreate(ctx *gin.Context) {

}

func Upload(ctx *gin.Context) {
	var (
		fileName string
	)
	errCode := 0
	defer func() {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		if errCode < 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    errCode,
				"message": "上传失败",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    errCode,
				"message": "上传成功",
			})
		}
	}()
	file, header, err := ctx.Request.FormFile("file")
	if header != nil {
		fileName = header.Filename
	}

	if err != nil {
		log.Printf("Failed to get form data, err:%s\n", err.Error())
		errCode = -1
		return
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Printf("Failed to get file data, err:%s\n", err.Error())
		errCode = -2
		return
	}

	_, _ = file.Seek(0, 0)
	storeType := yml_config.CreateYamlFactory().GetString("Store.CurrentStoreType")

	if storeType == consts.StoreCephStore {
		// 写入ceph存储
		baseRootDir := yml_config.CreateYamlFactory().GetString("Store.Ceph.CephRootDir")
		data, _ := ioutil.ReadAll(file)
		cephPath := baseRootDir + common.Sha1(buf.Bytes())
		_ = ceph.PutObject("upload", cephPath, data)
		// ToDO
		//ceph.GetCephBucket("upload")
		//ctx.JSON(http.StatusOK, gin.H{
		//	"name": fileName,
		//	"ext":  "jpeg", // TODO
		//	"url": fmt.Sprintf("http://%s/file/download?filehash=%s&username=%s&token=%s",
		//		c.Request.Host, filehash, username, token),
		//})
	} else if storeType == consts.StoreOssStore {
		// 写入oss存储
		baseRootDir := yml_config.CreateYamlFactory().GetString("Store.Oss.OSSRootDir")
		ossPath := baseRootDir + common.Sha1(buf.Bytes())
		err := oss.Bucket().PutObject(ossPath, file)
		if err != nil {
			log.Println(err.Error())
			errCode = -5
			return
		}
		signUrl := oss.DownloadURL(ossPath)
		ctx.JSON(http.StatusOK, gin.H{
			"name": fileName,
			"ext":  "jpeg", // TODO
			"url":  signUrl,
		})

	}
}

func ListGoodsSearchPriority(ctx *gin.Context) {

}

func UpdateGoodsSearchPriority(ctx *gin.Context) {

}

func ListGoodsSearchKeyWord(ctx *gin.Context) {
	queryParams := make(map[string]interface{})

	keyWords := ctx.DefaultQuery("keywords", "")
	pageNo, _ := strconv.Atoi(ctx.DefaultQuery("page_no", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))

	queryParams["page_no"] = pageNo
	queryParams["keywords"] = keyWords
	queryParams["page_size"] = pageSize

	data, dataTotal := model.CreateKeyWordSearchHistoryFactory("").List(queryParams)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       data,
		"data_total": dataTotal,
		"page_no":    pageNo,
		"page_size":  pageSize,
	})
}
