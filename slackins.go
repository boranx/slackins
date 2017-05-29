package main

import (
	"github.com/sbstjn/hanu"
	"log"
)

func main() {
	env := envOperations{}

	slack, err := hanu.New(env.exist("SLACK_BOT_API_TOKEN"))

	if err != nil {
		log.Fatal(err)
	}

	slack.Command("execute <job> <parameters>", func(conv hanu.ConversationInterface) {
		job, _ := conv.String("job")
		params, _ := conv.String("parameters")
		jenkins := Construct(env.exist("URI"), job, env.exist("TOKEN"), *parse(params))
		achieve := trigger(*jenkins)
		if !achieve {
			log.Fatalf("jenkins trigger job failed Expected : true, got : False")
		}

		conv.Reply("Job Execution Success : " + job)
	})

	slack.Listen()
}
