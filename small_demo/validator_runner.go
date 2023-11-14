package main

import (
	"context"
	"errors"
	"fmt"
)

type Validator interface {
	Validate(ctx context.Context) error
}

type Validator1 struct{}

func (v Validator1) Validate(ctx context.Context) error {
	fmt.Println("Validator1 is validating...")
	return nil
}

type Validator2 struct{}

func (v Validator2) Validate(ctx context.Context) error {
	fmt.Println("Validator2 is validating...")
	return errors.New("Validator2 validation failed")
}

type ValidatorRunner struct {
	validators []Validator
}

func (vr *ValidatorRunner) Run(ctx context.Context) error {
	for _, validator := range vr.validators {
		if err := validator.Validate(ctx); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	ctx := context.Background()
	runner := &ValidatorRunner{
		validators: []Validator{Validator1{}, Validator2{}},
	}

	if err := runner.Run(ctx); err != nil {
		fmt.Println("Validation failed:", err)
	} else {
		fmt.Println("All validations passed")
	}
}

/*
这种设计的优点：
扩展性：如果需要添加新的校验逻辑，只需要实现Validator接口，然后添加到validators切片中即可，无需修改已有的代码。
解耦：每个校验器都是独立的，它们之间没有依赖关系，可以独立进行测试和修改。

这种设计的缺点：
错误处理：在这个示例中，如果一个校验器校验失败，就会立即返回错误，后面的校验器不会被执行。这可能不是你想要的行为，你可能希望即使有校验器失败，也能执行所有的校验器，然后返回所有的错误。
性能：所有的校验器都是串行执行的，如果校验逻辑比较复杂，可能会影响性能。你可能需要使用并发或者其他优化手段来提高性能。

在面试时，你可以根据上述的内容来回答。
首先解释一下这种设计的基本思想和实现方式，然后讨论一下它的优点和缺点，最后可以提一下如何解决这些问题。这样可以展示出你对这个问题的深入理解和实践经验。

Q: 能够再封装一次吗，最后调用者只需要执行Run(validators)，就达到了校验的目的
A: 我们可以创建一个新的结构体，比如叫做ValidatorRunner，它包含一个Validator切片，并提供一个Run方法来执行所有的校验器。
这样，调用者只需要创建一个ValidatorRunner，然后调用它的Run方法即可

Q: 创建一个新的结构体这种思路非常巧妙，你是怎么知道的，从哪里学习到的呢
A: 这种思路是基于面向对象编程和设计模式的基本原则，特别是封装和抽象的原则。
在面向对象编程中，我们通常会把相关的数据和操作封装到一个对象中，这样可以提高代码的可读性和可维护性。
在设计模式中，我们通常会把复杂的操作抽象成一个接口或者一个抽象类，然后提供多种不同的实现，这样可以提高代码的灵活性和可扩展性。

这种思路并不是我个人独创的，而是计算机科学和软件工程领域的共识。
你可以在很多编程书籍和教程中学习到这些原则，比如《设计模式：可复用面向对象软件的基础》、《重构：改善既有代码的设计》等。
你也可以在很多开源项目中看到这些原则的应用，比如Go语言的标准库、Docker、Kubernetes等。

如果你想要深入学习这些原则和模式，我建议你阅读一些经典的编程书籍，比如我刚才提到的那两本。
你也可以参考一些优秀的开源项目，看看他们是如何设计和实现代码的。
同时，你也可以参加一些编程课程或者研讨会，和其他开发者交流学习经验。通过这些方式，你可以逐渐提高你的编程技能和设计能力。
*/
