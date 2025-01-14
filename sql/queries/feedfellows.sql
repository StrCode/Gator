-- name: CreateFeedFellow :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (id, user_id, feed_id, created_at, updated_at)
  VALUES ($1, $2, $3, $4, $5)
  RETURNING *
)
SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
  FROM inserted_feed_follow
  INNER JOIN users ON users.id = inserted_feed_follow.user_id
  INNER JOIN feeds ON feeds.id = inserted_feed_follow.feed_id;


-- name: GetFeedFollowsForUser :many
SELECT 
  feeds.name as FeedName,
  users.name as UserName
  FROM feed_follows
  INNER JOIN users ON feed_follows.user_id = users.id
  INNER JOIN feeds ON feed_follows.feed_id = feeds.id
  WHERE feed_follows.user_id = $1;
