package consts

const (
	WAIT     = "WAIT"
	UNDERWAY = "UNDERWAY"
	END      = "END"
)

const (
	/**
	 * 编辑中
	 */
	SeckillEditing = "EDITING"

	/**
	 * 已发布
	 */
	SeckillRelease = "RELEASE" // 已发布

	/**
	 * 已开启
	 */
	SeckillOpen = "OPEN" // 已开启

	/**
	 * 已关闭
	 */
	SeckillClosed = "CLOSED" // 已关闭
	/**
	 * 已结束
	 */
	SeckillOver = "OVER" // 已结束
)

var SeckillMap = map[string]string{
	SeckillEditing: "编辑中",
	SeckillRelease: "已发布",
	SeckillOpen:    "已开启",
	SeckillClosed:  "已关闭",
	SeckillOver:    "已结束",
}
