package main

import (
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	// 创建一个示例的Protocol Buffers消息
	message := &MyMessage{
		StringField: "Hello, World!",
		BoolField:   true,
		// Int32Field is not set
		// TimestampField is not set
	}
	// 定义MarshalOptions
	opts := protojson.MarshalOptions{
		EmitUnpopulated: true, // 这里设置为true，表示生成Go对象中的空值或未设置字段
	}
	// 将消息编码为JSON
	json, err := opts.Marshal(message)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// 打印生成的JSON
	fmt.Println(string(json))
}

// 可以通过`protojson.MarshalOptions`结构体中的`EmitUnpopulated`字段来控制是否生成Go对象中的空值或未设置字段。以下是一个示例代码：
// 在上面的示例中，`EmitUnpopulated` 被设置为 `true`，所以即使 `Int32Field` 和 `TimestampField` 没有被设置，它们也会在生成的JSON中出现。
/*
EmitUnpopulated 是 protojson.MarshalOptions 的一个选项，当设置为 true 时，它会在序列化 Protocol Buffers 消息为 JSON 时包含所有字段，即使这些字段的值是其类型的零值（例如，布尔类型的 false，数值类型的 0，字符串类型的空字符串，以及空的数组和对象）。
这个选项的主要好处是它提供了更完整的 JSON 输出，这在某些情况下可能是必需的。例如，如果你的 API 的用户期望在 JSON 响应中看到所有字段，即使这些字段的值是空的，那么你就需要使用 EmitUnpopulated 选项。
此外，这个选项也可以帮助你在调试时更容易地理解你的数据。如果你的 Protocol Buffers 消息有很多字段，而你只设置了其中的一部分，那么在查看 JSON 输出时，可能会很难确定哪些字段没有被设置。如果你使用了 EmitUnpopulated 选项，那么所有字段都会被包含在 JSON 输出中，这样你就可以很容易地看到哪些字段的值是空的。
然而，这个选项也有一些缺点。首先，它会增加 JSON 输出的大小，因为它包含了所有的字段，而不仅仅是那些被设置了值的字段。其次，它可能会导致你的 JSON 输出包含一些不必要的信息，这可能会使得你的 API 的用户更难理解你的数据。因此，你应该根据你的具体需求来决定是否使用 EmitUnpopulated 选项。
*/
