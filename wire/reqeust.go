package wire

type InfoType struct {
	IType string `form:"iType" json:"iType" binding:"required,oneof=cpu memory" label:"信息类型"`
}
