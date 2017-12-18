package mp

import (
	"encoding/json"
	"fmt"
	"github.com/BruceMaa/GoChat/wechat/common"
	"mime/multipart"
)

const (
	WECHAT_MEDIA_UPLOAD_API         = `https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s`          // 新增临时素材
	WECHAT_MEDIA_GET_API            = `https://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s`         // 获取临时素材
	WECHAT_MEDIA_UPLOADIMG_API      = `https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=%s`               // 上传图文消息内的图片获取URL
	WECHAT_MATERIAL_ADD_API         = `https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=%s&type=%s` // 新增其他类型永久素材
	WECHAT_MATERIAL_GET_API         = `https://api.weixin.qq.com/cgi-bin/material/get_material?access_token=%s`         // 获取永久素材
	WECHAT_MATERIAL_DELETE_API      = `https://api.weixin.qq.com/cgi-bin/material/del_material?access_token=%s`         // 删除永久素材
	WECHAT_MATERIAL_GET_COUNT_API   = `https://api.weixin.qq.com/cgi-bin/material/get_materialcount?access_token=%s`    // 获取素材总数
	WECHAT_MATERIAL_BATCHGET_API    = `https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=%s`    // 获取素材列表
	WECHAT_MATERIAL_NEWS_ADD_API    = `https://api.weixin.qq.com/cgi-bin/material/add_news?access_token=%s`             // 新增永久图文素材
	WECHAT_MATERIAL_NEWS_UPDATE_API = `https://api.weixin.qq.com/cgi-bin/material/update_news?access_token=%s`          // 修改永久图文素材
)

type WechatMediaType string

const (
	WECHAT_MEDIA_TYPE_IMAGE WechatMediaType = "image" // 素材类型：图片
	WECHAT_MEDIA_TYPE_VOICE WechatMediaType = "voice" // 素材类型：语音
	WECHAT_MEDIA_TYPE_VIDEO WechatMediaType = "video" // 素材类型：视频
	WECHAT_MEDIA_TYPE_THUMB WechatMediaType = "thumb" // 素材类型：缩略图，主要用于视频与音乐格式的缩略图
	WECHAT_MEDIA_TYPE_NEWS  WechatMediaType = "news"  // 素材类型：图文
)

type MediaResponse struct {
	MediaType WechatMediaType `json:"type"`      // 媒体文件类型，分别有图片（image）、语音（voice）、视频（video）和缩略图（thumb，主要用于视频与音乐格式的缩略图）
	MediaId   string          `json:"media_id"`  // 媒体文件上传后，获取标识
	CreateAt  int64           `json:"create_at"` // 媒体文件上传时间戳
	common.PublicError
}

// 图文消息，是否显示封面
type WechatArticleShowCoverPic int

// 图文消息，是否打开评论
type WechatArticleNeedOpenComment int

// 图文消息，是否只有粉丝才可以评论
type WechatArticleOnlyFansCanComment int

const (
	WECAHT_ARTICLE_SHOW_COVER_PIC_FALSE WechatArticleShowCoverPic = 0 // 不显示封面
	WECAHT_ARTICLE_SHOW_COVER_PIC_TRUE  WechatArticleShowCoverPic = 1 // 显示封面

	WECHAT_ARTICLE_NEED_OPEN_COMMENT_NO  WechatArticleNeedOpenComment = 0 // 不打开评论
	WECHAT_ARTICLE_NEED_OPEN_COMMENT_YES WechatArticleNeedOpenComment = 1 // 打开评论

	WECHAT_ARTICLE_ONLY_FANS_CAN_COMMENT_NO  WechatArticleOnlyFansCanComment = 0 // 否，所有人都可以评论
	WECHAT_ARTICLE_ONLY_FANS_CAN_COMMENT_YES WechatArticleOnlyFansCanComment = 1 // 是滴, 只有粉丝才能评论
)

type (
	Articles struct {
		Articles []Article `json:"articles"`
	}
	Article struct {
		Title              string                          `json:"title"`                           // 标题
		ThumbMediaId       string                          `json:"thumb_media_id"`                  // 图文消息的封面图片素材id（必须是永久mediaID）
		Author             string                          `json:"author,omitempty"`                // 作者
		Digest             string                          `json:"digest,omitempty"`                // 图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空。如果本字段为没有填写，则默认抓取正文前64个字。
		ShowCoverPic       WechatArticleShowCoverPic       `json:"show_cover_pic"`                  // 是否显示封面，0为false，即不显示，1为true，即显示
		Content            string                          `json:"content"`                         // 图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且此处会去除JS,涉及图片url必须来源 "上传图文消息内的图片获取URL"接口获取。外部图片url将被过滤。
		ContentSourceUrl   string                          `json:"content_source_url"`              // 图文消息的原文地址，即点击“阅读原文”后的URL
		Url                string                          `json:"url,omitempty"`                   // 图文页的URL
		NeedOpenComment    WechatArticleNeedOpenComment    `json:"need_open_comment,omitempty"`     // 是否打开评论，0不打开，1打开
		OnlyFansCanComment WechatArticleOnlyFansCanComment `json:"only_fans_can_comment,omitempty"` // 是否粉丝才可评论，0所有人可评论，1粉丝才可评论
	}
	ArticleNews struct {
		NewsItem []Article `json:"news_item"`
		common.PublicError
	}
	ArticleUpdate struct {
		MediaId  string  `json:"media_id"` // 要修改的图文消息的id
		Index    int     `json:"index"`    // 要更新的文章在图文消息中的位置（多图文消息时，此字段才有意义），第一篇为0
		Articles Article `json:"articles"`
	}
)

type (
	Material struct {
		MediaId string `json:"media_id"`      // 新增的图文消息素材的media_id
		Url     string `json:"url,omitempty"` // 上传图片的URL，可放置图文消息中使用。
		common.PublicError
	}

	MaterialVideo struct {
		Title        string `json:"title"`        // 视频素材的标题
		Introduction string `json:"introduction"` // 视频素材的描述
		DownUrl      string `json:"down_url"`     // 视频素材的地址
		common.PublicError
	}

	MaterialCount struct {
		VoiceCount int `json:"voice_count"` // 语音总数量
		VideoCount int `json:"video_count"` // 视频总数量
		ImageCount int `json:"image_count"` // 图片总数量
		NewsCount  int `json:"news_count"`  // 图文总数量
		common.PublicError
	}

	BatchMaterialRequest struct {
		Type   WechatMediaType `json:"type"`   // 素材的类型，图片（image）、视频（video）、语音 （voice）、图文（news）
		Offset int             `json:"offset"` // 从全部素材的该偏移位置开始返回，0表示从第一个素材 返回
		Count  int             `json:"count"`  // 返回素材的数量，取值在1到20之间
	}

	BatchMaterialNewsResponse struct {
		TotalCount int                 `json:"total_count"` // 该类型的素材的总数
		ItemCount  int                 `json:"item_count"`  // 本次调用获取的素材的数量
		Item       []BatchMaterialNews `json:"item"`
		common.PublicError
	}

	BatchMaterialNews struct {
		MediaId    string      `json:"media_id"`
		Content    ArticleNews `json:"content"`
		UpdateTime int64       `json:"update_time"` // 这篇图文消息素材的最后更新时间
	}

	BatchMaterialResponse struct {
		TotalCount int                 `json:"total_count"` // 该类型的素材的总数
		ItemCount  int                 `json:"item_count"`  // 本次调用获取的素材的数量
		Item       []BatchMaterialNews `json:"item"`
		common.PublicError
	}

	BatchMaterial struct {
		MediaId    string `json:"media_id"`
		Name       string `json:"name"` // 文件名称
		UpdateTime int64  `json:"update_time"`
		Url        string `json:"url"`
	}
)

const (
	upload_media_field_name       = "media"       // 上传素材时，文件入参名称
	add_material_video_field_desp = "description" // 添加视频素材时，必须添加视频说明
)

/*
1、临时素材media_id是可复用的。

2、媒体文件在微信后台保存时间为3天，即3天后media_id失效。

3、上传临时素材的格式、大小限制与公众平台官网一致。

图片（image）: 2M，支持PNG\JPEG\JPG\GIF格式

语音（voice）：2M，播放长度不超过60s，支持AMR\MP3格式

视频（video）：10MB，支持MP4格式

缩略图（thumb）：64KB，支持JPG格式

4、需使用https调用本接口。
*/
// 新增临时素材
func (wm *WechatMp) UploadMedia(accessToken string, mediaType WechatMediaType, file multipart.FileHeader) (*MediaResponse, error) {
	//TODO 根据文件类型判断文件大小
	data, err := file.Open()
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "UploadMedia file.Open() error: %+v\n", err)
		return nil, err
	}
	defer data.Close()
	formFile := &common.MultipartFormFile{
		FieldName: upload_media_field_name,
		FileName:  file.Filename,
		Reader:    data,
	}
	resp, err := common.HTTPPostForm(fmt.Sprintf(WECHAT_MEDIA_UPLOAD_API, accessToken, mediaType), nil, *formFile)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "UploadMedia common.HTTPUpload error: %+v\n", err)
		return nil, err
	}
	var mediaResponse MediaResponse
	if err = json.Unmarshal(resp, &mediaResponse); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "UploadMedia json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &mediaResponse, nil
}

// 获取临时素材
func (wm *WechatMp) GetMedia(accessToken, mediaId string) (*[]byte, *common.PublicError, error) {
	resp, err := common.HTTPGet(fmt.Sprintf(WECHAT_MEDIA_GET_API, accessToken, mediaId))
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "GetMedia http get error: %+v\n", err)
		return nil, nil, err
	}
	var result common.PublicError
	if err = json.Unmarshal(resp, &result); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "GetMedia json.Unmarshal error: %+v\n", err)
		return nil, nil, err
	}
	if result.Errcode != 0 {
		// 获取失败
		fmt.Fprintf(common.WechatErrorLoggerWriter, "GetMedia errcode: %d, errmsg: %s\n", result.Errcode, result.Errmsg)
		return nil, &result, nil
	}
	return &resp, nil, nil
}

// 新增永久图文素材
func (wm *WechatMp) AddMaterialNews(accessToken string, articles *Articles) (*Material, error) {
	resp, err := common.HTTPPostJson(fmt.Sprintf(WECHAT_MATERIAL_NEWS_ADD_API, accessToken), &articles)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "AddMaterialNews http post error: %+v\n", err)
		return nil, err
	}
	var material Material
	if err = json.Unmarshal(resp, &material); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "AddMaterialNews json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &material, nil
}

// 上传图文消息内的图片获取URL
func (wm *WechatMp) UploadMediaImg(accessToken string, file multipart.FileHeader) (*Material, error) {
	data, err := file.Open()
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "UploadMediaImg file.Open() error: %+v\n", err)
		return nil, err
	}
	defer data.Close()
	formFile := &common.MultipartFormFile{
		FieldName: upload_media_field_name,
		FileName:  file.Filename,
		Reader:    data,
	}
	resp, err := common.HTTPPostForm(fmt.Sprintf(WECHAT_MEDIA_UPLOADIMG_API, accessToken), nil, *formFile)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "UploadMediaImg common.HTTPPostForm error: %+v\n", err)
		return nil, err
	}
	var material Material
	if err = json.Unmarshal(resp, &material); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "UploadMediaImg json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &material, nil
}

// 新增其他类型永久素材
// 在上传视频素材时需要POST另一个表单，id为description，包含素材的描述信息，内容格式为JSON
//TODO 待验证
func (wm *WechatMp) AddMaterial(accessToken string, mediaType WechatMediaType, file multipart.FileHeader, videoRequest *MaterialVideo) (*Material, error) {
	if mediaType == WECHAT_MEDIA_TYPE_VIDEO && videoRequest == nil {
		return nil, fmt.Errorf("上传视频素材，需要填写视频信息MaterialVideoRequest")
	}

	data, err := file.Open()
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "AddMaterial file.Open() error: %+v\n", err)
		return nil, err
	}
	defer data.Close()
	formFile := &common.MultipartFormFile{
		FieldName: upload_media_field_name,
		FileName:  file.Filename,
		Reader:    data,
	}
	var fields = make(map[string]string)
	if videoRequest != nil {
		if data, err := json.Marshal(videoRequest); err != nil {
			return nil, err
		} else {
			fields[add_material_video_field_desp] = string(data)
		}
	}
	resp, err := common.HTTPPostForm(fmt.Sprintf(WECHAT_MATERIAL_ADD_API, accessToken, mediaType), fields, *formFile)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "AddMaterial common.HTTPPostForm error: %+v\n", err)
		return nil, err
	}
	var material Material
	if err = json.Unmarshal(resp, &material); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "AddMaterial json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &material, nil
}

// 获取永久图文素材
func (wm *WechatMp) GetNewsMaterial(accessToken, mediaId string) (*ArticleNews, error) {
	material := &Material{
		MediaId: mediaId,
	}
	resp, err := common.HTTPPostJson(fmt.Sprintf(WECHAT_MATERIAL_GET_API, accessToken), *material)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "GetNewsMaterial common.HTTPPostJson error: %+v\n", err)
		return nil, err
	}
	var articleNews ArticleNews
	if err = json.Unmarshal(resp, &articleNews); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "GetNewsMaterial json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &articleNews, nil
}

// 获取永久视频素材
func (wm *WechatMp) GetVideoMaterial(accessToken, mediaId string) (*MaterialVideo, error) {
	material := &Material{
		MediaId: mediaId,
	}
	resp, err := common.HTTPPostJson(fmt.Sprintf(WECHAT_MATERIAL_GET_API, accessToken), *material)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "GetVideoMaterial http post error: %+v\n", err)
		return nil, err
	}
	var materialVideo MaterialVideo
	if err = json.Unmarshal(resp, &materialVideo); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "GetVideoMaterial json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &materialVideo, nil
}

// 获取其他类型的素材，自己转换
func (wm *WechatMp) GetOtherMaterial(accessToken, mediaId string) ([]byte, error) {
	material := &Material{
		MediaId: mediaId,
	}
	return common.HTTPPostJson(fmt.Sprintf(WECHAT_MATERIAL_GET_API, accessToken), material)
}

// 删除永久素材
func (wm *WechatMp) DeleteMaterial(accessToken, mediaId string) (*common.PublicError, error) {
	material := &Material{
		MediaId: mediaId,
	}
	resp, err := common.HTTPPostJson(fmt.Sprintf(WECHAT_MATERIAL_DELETE_API, accessToken), *material)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "DeleteMaterial http post error: %+v\n", err)
		return nil, err
	}
	var result common.PublicError
	if err = json.Unmarshal(resp, &result); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "DeleteMaterial json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &result, nil
}

// 修改永久图文素材
func (wm *WechatMp) UpdateMaterial(accessToken string, article *ArticleUpdate) (*common.PublicError, error) {
	resp, err := common.HTTPPostJson(fmt.Sprintf(WECHAT_MATERIAL_NEWS_UPDATE_API, accessToken), &article)
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "UpdateMaterial http post error: %+v\n", err)
		return nil, err
	}
	var result common.PublicError
	if err = json.Unmarshal(resp, &result); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "Updatematerial json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &result, nil
}

// 获取素材总数
func (wm *WechatMp) GetMaterialCount(accessToken string) (*MaterialCount, error) {
	resp, err := common.HTTPGet(fmt.Sprintf(WECHAT_MATERIAL_GET_COUNT_API, accessToken))
	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "GetMaterialCount http get error: %+v\n", err)
		return nil, err
	}
	var materialCount MaterialCount
	if err = json.Unmarshal(resp, &materialCount); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "GetMaterialCount json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &materialCount, nil
}

// 批量获取图文素材列表
func (wm *WechatMp) BatchGetNewsMaterial(accessToken string, request *BatchMaterialRequest) (*BatchMaterialNewsResponse, error) {
	if request.Type != WECHAT_MEDIA_TYPE_NEWS {
		return nil, fmt.Errorf("此方法只能批量获取图文消息! 获取其他素材列表，请使用方法 BatchGetMaterial。\n")
	}

	if request.Count > 20 || request.Count < 1 {
		return nil, fmt.Errorf("获取素材的数据量需要在1到20之间!\n")
	}

	resp, err := common.HTTPPostJson(fmt.Sprintf(WECHAT_MATERIAL_BATCHGET_API, accessToken), &request)

	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "BatchGetNewsMaterial http post error: %+v\n", err)
		return nil, err
	}
	var batchMaterialNewsResp BatchMaterialNewsResponse
	if err = json.Unmarshal(resp, &batchMaterialNewsResp); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "BatchGetNewsMaterial json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &batchMaterialNewsResp, nil
}

// 批量获取素材列表，图文素材单独获取:BatchGetNewsMaterial
func (wm *WechatMp) BatchGetMaterial(accessToken string, request *BatchMaterialRequest) (*BatchMaterialResponse, error) {
	if request.Type == WECHAT_MEDIA_TYPE_NEWS {
		return nil, fmt.Errorf("此方法只能获取图片、语音、视频素材列表，若获取图文素材列表，请使用方法 BatchGetNewsMaterial。\n")
	}
	if request.Count > 20 || request.Count < 1 {
		return nil, fmt.Errorf("获取素材的数据量需要在1到20之间!\n")
	}
	resp, err := common.HTTPPostJson(fmt.Sprintf(WECHAT_MATERIAL_BATCHGET_API, accessToken), &request)

	if err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "BatchGetMaterial http post error: %+v\n", err)
		return nil, err
	}

	var batchMaterialResp BatchMaterialResponse
	if err = json.Unmarshal(resp, &batchMaterialResp); err != nil {
		fmt.Fprintf(common.WechatErrorLoggerWriter, "BatchGetMaterial json.Unmarshal error: %+v\n", err)
		return nil, err
	}
	return &batchMaterialResp, nil
}
