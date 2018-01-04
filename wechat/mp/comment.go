package mp

import (
	"fmt"
	"github.com/BruceMaa/Panda/wechat/common"
	"github.com/chanxuehong/wechat/json"
)

const (
	WechatCommentOpenApi        = `https://api.weixin.qq.com/cgi-bin/comment/open?access_token=%s`         // 打开已群发文章评论
	WechatCommentCloseApi       = `https://api.weixin.qq.com/cgi-bin/comment/close?access_token=%s`        // 关闭已群发文章评论
	WechatCommentListApi        = `https://api.weixin.qq.com/cgi-bin/comment/list?access_token=%s`         // 查看指定文章的评论数据
	WechatCommentMarkApi        = `https://api.weixin.qq.com/cgi-bin/comment/markelect?access_token=%s`    // 将评论标记精选
	WechatCommentUnmarkApi      = `https://api.weixin.qq.com/cgi-bin/comment/unmarkelect?access_token=%s`  // 将评论取消精选
	WechatCommentDeleteApi      = `https://api.weixin.qq.com/cgi-bin/comment/delete?access_token=%s`       // 删除评论
	WechatCommentReplyAddApi    = `https://api.weixin.qq.com/cgi-bin/comment/reply/add?access_token=%s`    // 回复评论
	WechatCommentReplyDeleteApi = `https://api.weixin.qq.com/cgi-bin/comment/reply/delete?access_token=%s` // 删除回复
)

// 评论类型
type WechatCommentType int

const (
	WechatCommentTypeAll      = 0 // 普通评论&精选评论
	WechatCommentTypeNormal   = 1 // 普通评论
	WechatCommentTypeFeatured = 2 // 精选评论
)

type (
	CommentRequest struct {
		MsgDataId     int               `json:"msg_data_id"`               // 群发返回的msg_data_id
		Index         int               `json:"index,omitempty"`           // 多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
		Begin         int               `json:"begin,omitempty"`           // 起始位置
		Count         int               `json:"count,omitempty"`           // 获取数目（>=50会被拒绝）
		Type          WechatCommentType `json:"type,omitempty"`            // type=0 普通评论&精选评论 type=1 普通评论 type=2 精选评论
		UserCommentId int               `json:"user_comment_id,omitempty"` // 用户评论id
		Content       string            `json:"content,omitempty"`         // 回复内容
	}

	CommentResponse struct {
		Total   int       `json:"total"` // 总数，非comment的size
		Comment []Comment `json:"comment"`
		common.PublicError
	}

	Comment struct {
		UserCommentId string            `json:"user_comment_id"` // 用户评论id
		Openid        string            `json:"openid"`          // 用户公众号openid
		CreateTime    int64             `json:"create_time"`     // 评论时间
		Content       string            `json:"content"`         // 评论内容
		CommentType   WechatCommentType `json:"comment_type"`    // /是否精选评论，0为false即非精选，1为true，即精选
		Reply         struct {
			Content    string `json:"content"`     // 作者回复内容
			CreateTime int64  `json:"create_time"` // 作者回复时间
		} `json:"reply"`
	}
)

// 打开已群发文章评论
// msgDataId:群发返回的msg_data_id, index: 多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
func (wm *WechatMp) OpenComment(accessToken string, msgDataId int, index int) (*common.PublicError, error) {
	comment := &CommentRequest{
		MsgDataId: msgDataId,
		Index:     index,
	}
	return pubComment(WechatCommentOpenApi, accessToken, comment)
}

// 关闭已群发文章评论
// msgDataId:群发返回的msg_data_id, index: 多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
func (wm *WechatMp) CloseComment(accessToken string, msgDataId int, index int) (*common.PublicError, error) {
	comment := &CommentRequest{
		MsgDataId: msgDataId,
		Index:     index,
	}
	return pubComment(WechatCommentCloseApi, accessToken, comment)
}

// 查看指定文章的评论数据
func (wm *WechatMp) ListComment(accessToken string, comment *CommentRequest) (*CommentResponse, error) {
	resp, err := common.HTTPPostJson(fmt.Sprintf(WechatCommentListApi, accessToken), &comment)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "ListComment http post error: %+v\n", err)
		return nil, err
	}

	var commentResponse CommentResponse
	if err = json.Unmarshal(resp, &commentResponse); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "ListComment json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &commentResponse, nil
}

// 将评论标记精选
func (wm *WechatMp) MarkComment(accessToken string, comment *CommentRequest) (*common.PublicError, error) {
	return pubComment(WechatCommentMarkApi, accessToken, comment)
}

// 将评论取消精选
func (wm *WechatMp) UnMarkComment(accessToken string, comment *CommentRequest) (*common.PublicError, error) {
	return pubComment(WechatCommentUnmarkApi, accessToken, comment)
}

// 删除评论
func (wm *WechatMp) DeleteComment(accessToken string, comment *CommentRequest) (*common.PublicError, error) {
	return pubComment(WechatCommentDeleteApi, accessToken, comment)
}

// 回复评论
func (wm *WechatMp) ReplyComment(accessToken string, comment *CommentRequest) (*common.PublicError, error) {
	return pubComment(WechatCommentReplyAddApi, accessToken, comment)
}

// 删除回复
func (wm *WechatMp) DeleteReply(accessToken string, comment *CommentRequest) (*common.PublicError, error) {
	return pubComment(WechatCommentReplyDeleteApi, accessToken, comment)
}

// 通用评论接口
func pubComment(url, accessToken string, comment *CommentRequest) (*common.PublicError, error) {
	resp, err := common.HTTPPostJson(fmt.Sprintf(url, accessToken), &comment)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "comment http post url:%s, error: %+v\n", url, err)
		return nil, err
	}
	var result common.PublicError
	if err = json.Unmarshal(resp, &result); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "comment json.Unmarshal url:%s, error: %+v\n", url, err)
		return nil, err
	}
	return &result, err
}
