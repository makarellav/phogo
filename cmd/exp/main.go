package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/makarellav/phogo/models"
	"log"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://user:password@localhost:1111/phogo")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close(context.Background())

	usersSrv := models.UserService{DB: conn}

	user, err := usersSrv.Create("test@test.com", "secret_password_123")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", user)
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	_, err = conn.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS users
	//(
	//    id         SERIAL PRIMARY KEY,
	//    first_name TEXT NOT NULL,
	//    last_name  TEXT NOT NULL,
	//    email      TEXT NOT NULL
	//);
	//
	//CREATE TABLE IF NOT EXISTS tweets
	//(
	//    id        SERIAL PRIMARY KEY,
	//    author_id INT  NOT NULL,
	//    text      TEXT NOT NULL
	//);
	//
	//CREATE TABLE IF NOT EXISTS replies
	//(
	//    id        SERIAL PRIMARY KEY,
	//    tweet_id  INT  NOT NULL,
	//    parent_id INT,
	//    text      TEXT NOT NULL
	//);
	//
	//CREATE TABLE IF NOT EXISTS likes
	//(
	//    user_id  INT NOT NULL,
	//    tweet_id INT NOT NULL
	//)
	//`)
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}

	//	for i := 0; i < 10; i++ {
	//		fullName := strings.Split(gofakeit.Name(), " ")
	//
	//		_, err = conn.Exec(context.Background(),
	//			`
	//INSERT INTO users (first_name, last_name, email)
	//VALUES ($1, $2, $3)
	//`, fullName[0], fullName[1], gofakeit.Email())
	//
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//	}

	//userID := 1
	//_, err = conn.Exec(context.Background(),
	//	`INSERT INTO tweets(author_id, text) VALUES ($1, $2)`,
	//	userID, gofakeit.Sentence(10))

	//if err != nil {
	//	log.Fatal(err)
	//}

	//
	//for i := 1; i <= 10; i++ {
	//	_, err := conn.Exec(context.Background(),
	//		`INSERT INTO likes (user_id, tweet_id) VALUES ($1, $2)`,
	//		i, tweetID)
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

	//var likes int
	//err = conn.QueryRow(context.Background(), `SELECT COUNT(*) FROM likes WHERE tweet_id = $1`, tweetID).Scan(&likes)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("Likes: %d", likes)

	//parentID := 2
	//
	//_, err = conn.Exec(context.Background(),
	//	`INSERT INTO replies(tweet_id, parent_id, text) VALUES ($1, $2, $3)`,
	//	tweetID, parentID, gofakeit.Sentence(10))
	//
	//if err != nil {
	//	log.Fatal(err)
	//}

	//	rows, _ := conn.Query(context.Background(), `SELECT *
	//FROM replies
	//WHERE tweet_id = $1
	//ORDER BY parent_id IS NULL DESC, parent_id;`, tweetID)
	//
	//	defer rows.Close()
	//
	//	reps, err := pgx.CollectRows(rows, pgx.RowToStructByName[Reply])
	//
	//	replyTree := getReplyTree(reps)
	//
	//	if rows.Err() != nil {
	//		log.Fatal(rows.Err())
	//	}
	//
	//	data, _ := json.Marshal(replyTree)
	//
	//	err = os.WriteFile("data.json", data, 0644)
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
}

//func getReplyTree(reps []Reply) RepliesMap {
//	replyTree := make(RepliesMap)
//
//	for _, reply := range reps {
//		parentID := reply.ParentID
//
//		if !parentID.Valid {
//			fmt.Println("top-level", reply)
//
//			replyTree[reply.ID] = reply
//		} else {
//			parent := replyTree[int(parentID.Int64)]
//
//			fmt.Println("parent", parent)
//			fmt.Println("child", reply)
//
//			if parent.Replies == nil {
//				parent.Replies = make(RepliesMap)
//			}
//
//			parent.Replies[reply.ID] = reply
//		}
//	}
//
//	return replyTree
//}
