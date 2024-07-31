package solidGarbanzo

import (
	"testing"
)

func TestExtendedMessages(t *testing.T)  {

	messages := map[string] string{
		"George":"Great to see you, George!", "Kyle":"Hi, Kyle. Welcome!",
		 "Tino":"Hi, Tino. Welcome!", "Tyler":"Great to see you, Tyler!",
	}
	messagesExtended, err := ExtendedMessages(messages)

	if(err != nil){
		t.Fatalf("an error: %v was received when nill was expected", err)
	}
	for name, extendedItem := range(messagesExtended){
		if(messages[name] == extendedItem || extendedItem == ""){
			t.Fatalf("The extendend message matches the orignal. original %v , exetnsion: %v",messages[name], extendedItem)
		}
	}

}

func TestEmptyExtendedMessages(t *testing.T)  {

	messages := map[string] string{
		"George":"Great to see you, George!", "Kyle":"Hi, Kyle. Welcome!",
		 "Tino":"Hi, Tino. Welcome!", "Tyler":"",
	}
	messagesExtended, err := ExtendedMessages(messages)

	if(err == nil || len(messagesExtended) > 0){
		t.Fatalf("an error:was exepcted but nill was returned or an empty message was extended")
	}	
}