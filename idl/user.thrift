namespace go user

struct douyin_user_register_request {
    1:required string username
    2:required string password
}

struct douyin_user_register_response {
    1:required i32 status_code
    2:optional string status_msg
    3:required i64 user_id
    4:required string token
}

struct douyin_user_login_request {
    1:required string username
    2:required string password
}

struct douyin_user_login_response {
    1:required i32 status_code
    2:optional string status_msg
    3:required i64 user_id
    4:required string token
}

struct douyin_user_request {
    1:required i64 user_id
    2:required string token
}

struct douyin_user_response {
    1:required i32 status_code
    2:optional string status_msg
    3:required User user
}

struct User {
    1:required i64 id
    2:required string name
    3:optional i64 follow_count
    4:optional i64 follower_count
    5:required bool is_follow
}

service UserService {
    douyin_user_register_response CreateUser(1:douyin_user_register_request req)
    douyin_user_login_response CheckUser(1:douyin_user_login_request req)
    douyin_user_response GetUser(1:douyin_user_request req)
}
