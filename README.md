# video-server
go语言实现的流媒体网站

api设计
用户
创建（注册）用户：URL:/user  Method:POST, SC:201 400,500
用户登录：URL: /user/:username  Method:POST, SC:200,400
获取用户基本信息：URL:/user/:username  Method:GET, SC:200,400,401,403,500
用户注销：URL:/user/:username  Method:DELETE, SC:204,400,401,403,500

用户资源
List all videos: URL:/user/:username/videos  Method:GET, SC:200,400,500
Get one videos: URL:/user/:username/videos/:vid-id   Method:GET, SC:200,400,500
Delete one video: URL:/user/:username/videos/:vid-id  Method:DELETE, SC:204,400,401,403,500

评论
Show comments: URL:/videos/:vid-id/comments  Method:GET, SC:200,400,500
Post a comment: URL:/videos/:vid-id/comments Method:POST, SC:201,400,500
Delete a comment: URL:/videos/:vid-id/comments/:comment-id  Method:DELETE, SC:204,400,401,403,500

handler->validation{1.request, 2.user}->business logic->response

db:
CREATE TABLE `users` (
`id` int UNSIGNED NOT NULL AUTO_INCREMENT,
`login_name` varchar(64) NOT NULL,
`pwd` text NOT NULL,
PRIMARY KEY (`id`) ,
UNIQUE INDEX `uq_login_name` (`login_name`)
);

CREATE TABLE `video_info` (
`id` varchar(64) NOT NULL,
`author_id` int UNSIGNED NOT NULL,
`name` text NOT NULL,
`display_ctime` text NOT NULL,
`create_time` datetime NOT NULL,
PRIMARY KEY (`id`) 
);

CREATE TABLE `comments` (
`id` varchar(64) NOT NULL,
`video_id` varchar(64) NOT NULL,
`author_id` int UNSIGNED NOT NULL,
`content` text NOT NULL,
`time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`) 
);

CREATE TABLE `sessions` (
`session_id` varchar(64) NOT NULL,
`TTL` tinytext NULL,
`login_name` varchar(64) NOT NULL,
PRIMARY KEY (`session_id`) 
);