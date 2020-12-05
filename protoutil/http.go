package protoutil

import (
	"io"
	"io/ioutil"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ReadJSONMessage reads the JSON from r and converts it into Protobuf Message.
func ReadJSONMessage(r io.Reader, msg proto.Message) error {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return protojson.Unmarshal(buf, msg)
}

// WriteJSONMessage converts Protobuf Message to JSON and write to w.
func WriteJSONMessage(w io.Writer, msg proto.Message) error {
	buf, err := protojson.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = w.Write(buf)
	return err
}
