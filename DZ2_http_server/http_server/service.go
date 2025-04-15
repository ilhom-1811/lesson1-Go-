package http_server

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	ErrorGetComents       = errors.New("ошибка при получении комментариев")
	ErrorReadBody         = errors.New("ошибка при чтении тела ответа")
	ErrorGetComment       = errors.New("ошибка при получении комментария")
	ErrorMarshalComment   = errors.New("ошибка при маршаллинге комментария")
	ErrorCreateComment    = errors.New("ошибка при создании комментария")
	ErrorUpdateComment    = errors.New("ошибка при обновлении комментария")
	ErrorDeleteComment    = errors.New("ошибка при удалении комментария")
	ErrorUnmarshalComment = errors.New("ошибка при разборе комментария")
)

const URL_C = "https://jsonplaceholder.typicode.com/comments"

type CommentService struct{}

func (service *CommentService) GetComments() ([]Comment, error) {
	resp, err := http.Get(URL_C)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrorGetComents, err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrorReadBody
	}

	var comments []Comment
	err = json.Unmarshal(body, &comments)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrorUnmarshalComment, err)
	}

	return comments, nil
}

func (service *CommentService) GetComment(id int) (*Comment, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%d", URL_C, id))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrorGetComment, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrorReadBody
	}

	var comment Comment
	err = json.Unmarshal(body, &comment)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrorUnmarshalComment, err)
	}

	return &comment, nil
}

func (service *CommentService) CreateComment(comment Comment) (*Comment, error) {
	commentData, err := json.Marshal(comment)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrorMarshalComment, err)
	}

	resp, err := http.Post(URL_C, "application/json", bytes.NewBuffer(commentData))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrorCreateComment, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrorReadBody
	}

	var createdComment Comment
	err = json.Unmarshal(body, &createdComment)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrorUnmarshalComment, err)
	}

	return &createdComment, nil
}

func (service *CommentService) UpdateComment(id int, updatedComment Comment) (*Comment, error) {
	commentData, err := json.Marshal(updatedComment)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrorMarshalComment, err)
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%d", URL_C, id), bytes.NewBuffer(commentData))
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании PUT-запроса: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrorUpdateComment, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrorReadBody
	}

	var comment Comment
	err = json.Unmarshal(body, &comment)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrorUnmarshalComment, err)
	}

	return &comment, nil
}

func (service *CommentService) DeletePost(id int) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%d", URL_C, id), nil)
	if err != nil {
		return fmt.Errorf("ошибка при создании DELETE-запроса: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrorDeleteComment, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	return fmt.Errorf("ошибка при удалении комментарии: %s", resp.Status)
}
