package explorer

import (
	"context"
	"github.com/HFO4/cloudreve/pkg/filesystem"
	"github.com/HFO4/cloudreve/pkg/serializer"
	"github.com/gin-gonic/gin"
)

// ItemMoveService 处理多文件/目录移动
type ItemMoveService struct {
	SrcDir string      `json:"src_dir" binding:"required,min=1,max=65535"`
	Src    ItemService `json:"src" binding:"exists"`
	Dst    string      `json:"dst" binding:"required,min=1,max=65535"`
}

// ItemService 处理多文件/目录相关服务
type ItemService struct {
	Items []string `json:"items" binding:"exists,dive,ne=/"`
	Dirs  []string `json:"dirs" binding:"exists,dive,ne=/"`
}

// Delete 删除对象
func (service *ItemService) Delete(ctx context.Context, c *gin.Context) serializer.Response {
	// 创建文件系统
	fs, err := filesystem.NewFileSystemFromContext(c)
	if err != nil {
		return serializer.Err(serializer.CodePolicyNotAllowed, err.Error(), err)
	}

	// 删除对象
	err = fs.Delete(ctx, service.Dirs, service.Items)
	if err != nil {
		return serializer.Err(serializer.CodeNotSet, err.Error(), err)
	}

	return serializer.Response{
		Code: 0,
	}

}

// Move 移动对象
func (service *ItemMoveService) Move(ctx context.Context, c *gin.Context) serializer.Response {
	// 创建文件系统
	fs, err := filesystem.NewFileSystemFromContext(c)
	if err != nil {
		return serializer.Err(serializer.CodePolicyNotAllowed, err.Error(), err)
	}

	// 移动对象
	err = fs.Move(ctx, service.Src.Dirs, service.Src.Items, service.SrcDir, service.Dst)
	if err != nil {
		return serializer.Err(serializer.CodeNotSet, err.Error(), err)
	}

	return serializer.Response{
		Code: 0,
	}

}

// Copy 复制对象
func (service *ItemMoveService) Copy(ctx context.Context, c *gin.Context) serializer.Response {
	// 复制操作只能对一个目录或文件对象进行操作
	if len(service.Src.Items)+len(service.Src.Dirs) > 1 {
		return serializer.ParamErr("只能复制一个对象", nil)
	}

	// 创建文件系统
	fs, err := filesystem.NewFileSystemFromContext(c)
	if err != nil {
		return serializer.Err(serializer.CodePolicyNotAllowed, err.Error(), err)
	}

	// 移动对象
	err = fs.Copy(ctx, service.Src.Dirs, service.Src.Items, service.SrcDir, service.Dst)
	if err != nil {
		return serializer.Err(serializer.CodeNotSet, err.Error(), err)
	}

	return serializer.Response{
		Code: 0,
	}

}
