# DB設計

- ユーザー
- 企業
- 記事
- タグ
- コメント
- ブログタグ
- ストック記事
- フォロータグ
- ユーザーフォロー(中間テーブル)
- 企業フォロー(中間テーブル)
- いいね(中間テーブル)

### ユーザー(users)
- id int PRIMARY_KEY
- name string
- email string
- password string
- nickname string
- description string
- hp_url string
- location string
- github_account string
- organization_id int
- is_deleted boolean
- created_at timestamp

### 企業(organizations)
- id int PRIMARY_KEY
- name string
- description string
- hp_url string
- location string
- email string
- is_deleted boolean
### 記事(posts)
- id int PRIMARY_KEY
- title string
- body string
- posted_by int(relation user_id)
- organization_id int
- created_at timestamp
- updated_at timestamp
- is_draft boolean
- is_deleted boolean

### タグ(tags)
- id int PRIMARY_KEY
- name string
- user_id int (relation user_id)
- created_at timestamp

### コメント(comments)
- id int PRIMARY_KEY
- content string
- post_id int (relation post_id)
- user_id int (relation user_id)
- created_at timestamp
- updated_at timestamp

### ブログタグ(blog_tags)
- id int PRIMARY_KEY
- post_id int (relation post_id)
- tag_id int (relation tag_id)

### ストック記事(stock_posts)
- id int PRIMARY_KEY
- post_id int (relation post_id)
- user_id int (relation user_id)
- created_at timestamp

### フォロータグ(follow_tags)
- id int PRIMARY_KEY
- tag_id int (relation tag_id)
- user_id int (relation user_id)

### ユーザーフォロー(user_follows)
- id int PRIMARY_KEY
- follow_user_id int (relation user_id)
- followed_user_id int (relation user_id)
- created_at timestamp

### 企業フォロー(organization_follows)
- id int PRIMARY_KEY
- follow_user_id int (relation user_id)
- followed_organization_id int (relation organization_id)
- created_at timestamp

### いいね(likes)
- id int PRIMARY_KEY
- like_user_id int (relation user_id)
- liked_user_id int (relation user_id)
