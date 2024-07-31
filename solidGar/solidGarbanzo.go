package solidGarbanzo

import (
	"errors"
	"fmt"
	"math/rand"
)

// Extends the greetings with a follow up question
func ExtendedMessages(messages map[string]string) (map[string]string, error)  {
	extended := map[string]string{}
	for key, message := range(messages){
		if(message == ""){
			return nil, errors.New("the message is empty")
		}
		extended[key] = fmt.Sprintf("%v %v",message, randomExtensions() ) 
	}
	return extended, nil
}

func randomExtensions() string {

	extensions := []string{
		"How are you doing?", "How was your day?", "What have you been up to?", "How has your week been going",
	}
	
	return extensions[rand.Intn(len(extensions))]
}