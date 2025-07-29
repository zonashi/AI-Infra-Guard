package websocket

import (
	"errors"
	"fmt"
	"github.com/Tencent/AI-Infra-Guard/internal/mcp"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func HandleList(root string, loadFile func(filePath string) (interface{}, error)) gin.HandlerFunc {
	var allItems []interface{}
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil // 忽略错误
		}
		if !d.IsDir() {
			item, err := loadFile(path)
			if err != nil {
				return err
			}
			allItems = append(allItems, item)
		}
		return nil
	})
	return func(c *gin.Context) {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  1,
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  0,
			"message": "success",
			"data": gin.H{
				"total": len(allItems),
				"items": allItems,
			},
		})
	}
}
func HandleCreate(readAndSave func(content string) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		type UploadRequest struct {
			FileContent string `json:"file_content" binding:"required"`
		}
		var eval UploadRequest
		if err := c.ShouldBindJSON(&eval); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 1, "message": "参数解析失败"})
			return
		}
		if err := readAndSave(eval.FileContent); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": 1, "message": "保存失败: " + err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": 0, "message": "创建评测集成功"})
	}
}

// HandleEdit 返回处理编辑请求的HandlerFunc
func HandleEdit(updateFunc func(name string, content string) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		if name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"status": 1, "message": "名称不能为空"})
			return
		}

		type EditRequest struct {
			FileContent string `json:"file_content" binding:"required"`
		}
		var req EditRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 1, "message": "参数解析失败"})
			return
		}

		if err := updateFunc(name, req.FileContent); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": 1, "message": "更新失败: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": 0, "message": "更新成功"})
	}
}

// HandleDelete 返回处理删除请求的HandlerFunc
func HandleDelete(deleteFunc func(name string) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		if name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"status": 1, "message": "名称不能为空"})
			return
		}

		if err := deleteFunc(name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": 1, "message": "删除失败: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": 0, "message": "删除成功"})
	}
}

const MCPROOT = "data/mcp"

func McpLoadFile(filePath string) (interface{}, error) {
	if filePath == "" {
		return nil, nil
	}
	if !strings.HasSuffix(filePath, ".yaml") {
		return nil, nil
	}
	return mcp.NewYAMLPlugin(filePath)
}

func mcpReadAndSave(content string) error {
	var config mcp.PluginConfig
	err := yaml.Unmarshal([]byte(content), &config)
	if err != nil {
		return err
	}
	if config.Info.ID == "" {
		return errors.New("config info id is empty")
	}
	if strings.Contains(config.Info.ID, "..") {
		return errors.New("config info id contains ..")
	}
	if config.Info.Name == "" {
		return errors.New("config info name is empty")
	}
	if config.PromptTemplate == "" {
		return errors.New("config prompt_template is empty")
	}
	// save
	filename := filepath.Join(MCPROOT, config.Info.ID+".yaml")
	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		return err
	}
	return nil
}

func mcpUpdateFunc(name string, content string) error {
	// 解析新的配置
	var config mcp.PluginConfig
	if err := yaml.Unmarshal([]byte(content), &config); err != nil {
		return err
	}
	if config.Info.ID == "" {
		return errors.New("config info id is empty")
	}
	if strings.Contains(config.Info.ID, "..") {
		return errors.New("config info id contains ..")
	}
	if config.Info.Name == "" {
		return errors.New("config info name is empty")
	}
	if config.PromptTemplate == "" {
		return errors.New("config prompt_template is empty")
	}

	// 构建文件路径
	oldPath := filepath.Join(MCPROOT, name+".yaml")
	newPath := filepath.Join(MCPROOT, config.Info.ID+".yaml")

	// 检查原文件是否存在
	if _, err := os.Stat(oldPath); os.IsNotExist(err) {
		return errors.New("original config file does not exist")
	}
	// 写入新内容
	if err := os.WriteFile(newPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}
	return nil
}

func mcpDeleteFunc(name string) error {
	// 构建文件路径
	filePath := filepath.Join(MCPROOT, name+".yaml")

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return errors.New("config file does not exist")
	}

	// 删除文件
	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("failed to delete config file: %w", err)
	}

	return nil
}
