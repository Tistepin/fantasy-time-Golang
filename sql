create table ft_area
(
    id            bigint auto_increment comment 'id'
        primary key,
    name          char(50)                           null comment '国家名字',
    delete_status bigint   default 1                 not null comment '逻辑删除状态 0-已删除 1-未删除',
    create_time   datetime default CURRENT_TIMESTAMP null comment '创建时间',
    edit_time     datetime                           null comment '修改时间'
)
    comment '国家区域';

create table ft_cartoon_works_details
(
    id                   bigint auto_increment comment 'id'
        primary key,
    works_id             bigint                             null comment '作品id',
    cartoon_chapter_id   bigint                             null comment '漫画章节ID 第几话',
    cartoon_chapter_name varchar(64)                        null comment '漫画章节名字',
    cartoon_pages        char(20)                           null comment '漫画页数',
    create_time          datetime default CURRENT_TIMESTAMP null comment '注册时间',
    edit_time            datetime                           null comment '修改时间',
    delete_status        bigint   default 1                 not null comment '逻辑删除状态 0-已删除 1-未删除',
    review_status        bigint   default 0                 null comment '审核状态 0-审核中 1-审核成功 2-审核失败'
)
    comment '漫画作品章节数量';

create table ft_category
(
    cat_id        bigint auto_increment comment '分类id'
        primary key,
    name          char(50)         null comment '分类名称',
    sort          int              null comment '排序',
    works_type    tinyint          null comment '作品分类 1-漫画 2-小说',
    icon          char(255)        null comment '图标地址',
    delete_status bigint default 1 not null comment '逻辑删除状态 0-已删除 1-未删除'
)
    comment '作品类型';

create table ft_contact
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime(3)     null,
    updated_at datetime(3)     null,
    deleted_at datetime(3)     null,
    owner_id   bigint unsigned null,
    target_id  bigint unsigned null,
    type       bigint          null,
    desc_info  longtext        null
)
    comment '好友关系表';

create index idx_ft_contact_deleted_at
    on ft_contact (deleted_at);

create table ft_message
(
    id          bigint unsigned auto_increment
        primary key,
    created_at  datetime(3) null,
    updated_at  datetime(3) null,
    deleted_at  datetime(3) null,
    user_id     bigint      not null comment '发送者ID',
    target_id   bigint      not null comment '接受者ID',
    type        bigint      not null comment '发送类型',
    media       bigint      not null comment '消息类型',
    content     longtext    null comment '消息内存',
    create_time datetime(3) not null comment '创建时间',
    read_time   datetime(3) not null comment '读取时间',
    pic         longtext    null comment '图片',
    url         longtext    null comment '浏览地址',
    `desc`      longtext    null comment '备注',
    amount      bigint      null comment '其他数字统计'
)
    comment '消息表';

create index idx_ft_message_deleted_at
    on ft_message (deleted_at);

create table ft_novel_works_details
(
    id                 bigint auto_increment comment 'id'
        primary key,
    works_id           bigint                             null comment '作品id',
    novel_chapter_id   char(20)                           null comment '小说章节ID 第几章',
    novel_chapter_name varchar(64)                        null comment '小说章节名字',
    novel_pages        char(20)                           null comment '漫画页数',
    create_time        datetime default CURRENT_TIMESTAMP null comment '注册时间',
    edit_time          datetime                           null comment '修改时间',
    delete_status      bigint   default 1                 not null comment '逻辑删除状态 0-已删除 1-未删除'
)
    comment '小说作品章节数量';

create table ft_popularity
(
    id                          bigint auto_increment comment 'id'
        primary key,
    works_id                    bigint                             null comment '作品id',
    name                        char(50)                           null comment '作品名称',
    works_type                  tinyint                            null comment '作品类型 1-漫画 2-小说 3-插图',
    works_status                tinyint                            null comment '作品状态 1-更新中 2-完结',
    works_popularity_count      bigint                             null comment '总人气 观看十章 加一',
    works_popularity_today      bigint                             null comment '今天的人气 每个用户每天看一章加一 第二天清零',
    works_popularity_three_days bigint                             null comment '三天的人气 每个用户每天看一章加一 第四天天清零',
    works_popularity_thisWeek   bigint                             null comment '本周的人气 每个用户每天看一章加一 下一周清零',
    works_popularity_thisMonth  bigint                             null comment '本月的人气 每个用户每天看一章加一 下一月清零',
    delete_status               bigint   default 1                 null comment '逻辑删除 0-已删除 1-未删除',
    create_time                 datetime default CURRENT_TIMESTAMP not null comment '注册时间',
    edit_time                   datetime default CURRENT_TIMESTAMP null comment '修改时间'
)
    comment '作品人气';

create table ft_role_user
(
    uid           bigint                             null comment '用户ID',
    rid           bigint                             null comment '角色ID',
    create_time   datetime default CURRENT_TIMESTAMP null comment '注册时间',
    edit_time     datetime                           null comment '修改时间',
    delete_status bigint   default 1                 not null comment '逻辑删除状态 0-已删除 1-未删除'
)
    comment '用户角色关联表';

create table ft_roles
(
    id            bigint auto_increment comment 'id'
        primary key,
    role_name     varchar(50)                        not null comment '角色名称',
    role          varchar(50)                        not null comment '角色',
    create_time   datetime default CURRENT_TIMESTAMP null comment '注册时间',
    edit_time     datetime                           null comment '修改时间',
    delete_status bigint   default 1                 not null comment '逻辑删除状态 0-已删除 1-未删除',
    constraint idx_user_role
        unique (role_name, role)
)
    comment '角色信息表';

create table ft_user
(
    id            bigint auto_increment comment 'id'
        primary key,
    username      char(64)                           null comment '用户名',
    password      varchar(64)                        null comment '密码',
    nickname      varchar(64)                        null comment '昵称',
    phone         varchar(20)                        null comment '手机号码',
    email         varchar(64)                        null comment '邮箱',
    header        varchar(500)                       null comment '头像',
    gender        tinyint                            null comment '性别',
    birth         date                               null comment '生日',
    city          varchar(500)                       null comment '所在城市',
    job           varchar(255)                       null comment '职业',
    sign          varchar(255)                       null comment '个性签名',
    source_type   tinyint                            null comment '用户来源',
    status        tinyint                            null comment '启用状态',
    create_time   datetime default CURRENT_TIMESTAMP null comment '注册时间',
    edit_time     datetime default CURRENT_TIMESTAMP null comment '修改时间',
    delete_status bigint   default 1                 not null comment '逻辑删除状态 0-已删除 1-未删除',
    fans          int      default 0                 null comment '粉丝',
    care          int      default 0                 null comment '关注人数',
    likes         int      default 0                 null comment '获得的点赞数量',
    constraint email
        unique (email),
    constraint phone
        unique (phone)
)
    comment '用户信息表';

create table ft_works
(
    works_id          bigint auto_increment comment 'id'
        primary key,
    works_name        varchar(64)                        null comment '作品名',
    default_image     varchar(128)                       null comment '作品默认展示图片',
    creator           varchar(64)                        null comment '创建人',
    works_creator     varchar(64)                        null comment '作品创作者',
    works_create_time datetime                           null comment '作品创作时间',
    works_area        char(2)                            null comment '作品创作地区',
    works_type        bigint                             null comment '作品类型 1-漫画 2-小说 3-插图',
    works_score       float                              null comment '作品评分 10分满分',
    works_renew       bigint                             null comment '作品更新至多少',
    works_popularity  bigint   default 0                 null comment '作品人气 用户阅读10章以上为1',
    works_describe    varchar(64)                        null comment '作品描述',
    works_category    varchar(64)                        null comment '作品分类',
    status            tinyint                            null comment '作品状态 1-可以看 2-不可以看',
    create_time       datetime default CURRENT_TIMESTAMP null comment '注册时间',
    edit_time         datetime default CURRENT_TIMESTAMP null comment '修改时间',
    delete_status     bigint   default 1                 not null comment '逻辑删除状态 0-已删除 1-未删除',
    works_status      tinyint                            null comment '作品状态 1-更新中 2-完结',
    review_status     bigint   default 0                 null comment 'review_status审核状态 0-审核中 1-审核成功 2-审核失败',
    constraint works_name
        unique (works_name)
)
    comment '作品信息';

create table ft_works_bookshelf
(
    id                bigint auto_increment comment 'id'
        primary key,
    user_id           bigint                             not null comment '用户ID',
    works_id          bigint                             not null comment '作品id',
    works_name        varchar(64)                        null comment '作品名',
    default_image     varchar(128)                       null comment '作品默认展示图片',
    works_renew       varchar(64)                        null comment '作品更新至多少',
    works_type        bigint                             null comment '作品类型 1-漫画 2-小说',
    works_creator     varchar(64)                        null comment '作品创作者',
    works_update_time datetime                           null comment '作品更新时间',
    delete_status     bigint   default 1                 not null comment '逻辑删除状态 0-删除 1-存在',
    create_time       datetime default CURRENT_TIMESTAMP null comment '注册时间',
    edit_time         datetime                           null comment '修改时间'
)
    comment '书架';

create table ft_works_chapter_detailed_viewing_content
(
    id                     bigint auto_increment comment 'id'
        primary key,
    works_chapter_id       bigint                             null comment '章节ID',
    user_id                bigint                             not null comment '用户ID',
    works_id               bigint                             not null comment '作品id',
    image_id               bigint                             null comment '该画画作品的该章节的第几个图片',
    review_status          bigint   default 1                 not null comment '审核状态 0-审核中 1-审核成功 2-审核失败',
    delete_status          bigint   default 1                 not null comment '逻辑删除状态 0-删除 1-存在',
    create_time            datetime default CURRENT_TIMESTAMP null comment '注册时间',
    edit_time              datetime                           null comment '修改时间',
    works_chapter_Location varchar(128)                       not null comment '章节数据存储位置'
)
    comment '作品章节详细观看内容 例如小说第几章位置,漫画第一话的第一张图';

create table ft_works_default_image
(
    id                  bigint auto_increment comment 'id'
        primary key,
    works_id            bigint                             not null comment '作品id',
    works_default_image varchar(128)                       not null comment '图片服务请求数据位置',
    review_status       bigint   default 1                 not null comment '审核状态 0-审核中 1-审核成功 2-审核失败',
    delete_status       bigint   default 1                 not null comment '逻辑删除状态 0-删除 1-存在',
    create_time         datetime default CURRENT_TIMESTAMP null comment '注册时间',
    edit_time           datetime                           null comment '修改时间'
)
    comment '作品封面图片服务请求路径存储';

create table ft_works_upload
(
    id            bigint auto_increment comment 'id'
        primary key,
    user_id       bigint                             not null comment '用户ID',
    works_id      bigint                             not null comment '作品id',
    review_status bigint   default 1                 not null comment '审核状态 0-审核中 1-审核成功 2-审核失败',
    delete_status bigint   default 1                 not null comment '逻辑删除状态 0-删除 1-存在',
    create_time   datetime default CURRENT_TIMESTAMP null comment '注册时间',
    edit_time     datetime                           null comment '修改时间'
)
    comment '作品上传信息';

create table ft_works_watch_history
(
    id                                  bigint auto_increment comment 'id'
        primary key,
    user_id                             bigint                             not null comment '用户ID',
    works_id                            bigint                             not null comment '作品id',
    works_name                          varchar(64)                        null comment '作品名',
    default_image                       varchar(128)                       null comment '作品默认展示图片',
    works_history_viewing_chapter       bigint                             null comment '用户观看到第几章',
    works_history_viewing_chapter_id    bigint                             null comment '用户观看到第几章的ID',
    works_history_viewing_chapter_image bigint                             null comment '用户观看到第几章的的第几张图片',
    works_type                          tinyint                            null comment '作品分类 1-漫画 2-小说',
    works_status                        tinyint                            null comment '作品状态 1-更新中 2-完结',
    delete_status                       bigint   default 1                 not null comment '逻辑删除状态 0-删除 1-存在',
    create_time                         datetime default CURRENT_TIMESTAMP null comment '注册时间',
    edit_time                           datetime default CURRENT_TIMESTAMP null comment '修改时间'
)
    comment '观看历史记录';

create
    definer = root@`%` function GET_AREA_NAME(AREA_id varchar(20)) returns varchar(20) reads sql data
begin
    declare cid varchar(20);      #在函数中定义一个变量，用来接收函数返回值
    select ft_area.name into cid   #把查询结果赋值给cid变量
    from ft_area
    where ft_area.id=AREA_id;
    return cid;                   #函数返回值，返回cid
end;

create
    definer = root@`%` function GET_CATEGORY_NAME(CATEGORY_ID varchar(20)) returns varchar(20)
begin
    declare CATEGORY_NAME varchar(20);
    SELECT group_concat(name) into CATEGORY_NAME
    FROM ft_category
    WHERE FIND_IN_SET(cat_id, CATEGORY_ID);
    return CATEGORY_NAME;
end;

