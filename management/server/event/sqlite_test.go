package event

import (
	"fmt"
	"testing"
	"time"
)

func TestNewSQLiteStore(t *testing.T) {
	dataDir := t.TempDir()
	store, err := NewSQLiteStore(dataDir)
	if err != nil {
		t.Fatal(err)
		return
	}

	accountID := "account_1"

	for i := 0; i < 10; i++ {
		_, err = store.Save(Event{
			Timestamp:  time.Now(),
			Operation:  "cool_operation_" + fmt.Sprint(i),
			Type:       ManagementEvent,
			ModifierID: "user_" + fmt.Sprint(i),
			TargetID:   "peer_" + fmt.Sprint(i),
			AccountID:  accountID,
		})
		if err != nil {
			t.Fatal(err)
			return
		}
	}

	result, err := store.GetSince(accountID, time.Now().Add(-30*time.Second))
	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Println(len(result))

	result, err = store.GetLast(accountID, 5)
	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Println(len(result))
}