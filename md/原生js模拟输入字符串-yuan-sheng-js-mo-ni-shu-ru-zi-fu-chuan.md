---
title: 原生js模拟输入字符串
date: "2023-05-22 17:29:52"
updated: "2023-05-22 17:29:52"
url: https://p00q.cn/?p=874
categories:
    - 瞎折腾
    - 奇技淫巧
    - 开发
tags:
    - html
    - 模拟输入
summary: |-
    这个例子模拟了逐个字符输入文本的过程。首先获取了一个文本输入框元素，并定义了要输入的字符串和当前输入字符的索引。然后定义了一个 `simulateInput` 函数用于模拟逐个字符输入。在函数内部，检查是否还有字符需要输入，如果有，则获取当前要输入的字符，并创建一个 `InputEvent` 对象，将字符添加到文本输入框中，并分派 `inputEvent` 事件。然后增加当前输入字符的索引，并使用 `setTimeout` 函数在 100 毫秒后模拟下一个字符的输入。最后调用 `simulateInput` 函数开始模拟输入。

    整个过程模拟了文本逐个字符输入的效果。可以根据需要调整输入速度，通过调整 `setTimeout` 函数的延迟时间来实现。
id: "874"
---

在这个例子中，我们首先获取了一个文本输入框元素，并定义了要输入的字符串和当前输入字符的索引。然后，我们定义了一个 `simulateInput` 函数，该函数用于模拟逐个字符输入。在 `simulateInput` 函数中，我们首先检查是否还有字符需要输入，如果有，就获取当前要输入的字符，并创建一个 `InputEvent` 对象，将当前字符添加到文本输入框中，并将 `inputEvent` 分派到该元素上。然后，我们增加当前输入字符的索引，并使用 `setTimeout` 函数在 100 毫秒后模拟下一个字符的输入。最后，我们调用 `simulateInput` 函数开始模拟输入。

这个例子中的输入速度是固定的，每个字符之间间隔 100 毫秒。如果需要模拟不同的输入速度，可以根据需要调整 `setTimeout` 函数的延迟时间。

```
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>

<body>
    <input id="test-i" class="iii" type="text" />
    <script src="https://cdn.bootcdn.net/ajax/libs/jquery/3.6.4/jquery.min.js"></script>
</body>
<script>

    $("#test-i").on("input propertychange", function () {
        console.log($("#test-i").val())
    })

    function test() {
        const inputElement = document.getElementById('test-i'); // 获取文本输入框元素
        const inputString = 'Hello, world!'; // 要输入的字符串
        let index = 0; // 当前输入字符的索引
        function simulateInput() {
            if (index < inputString.length) { // 如果还有字符需要输入
                const char = inputString.charAt(index); // 获取当前要输入的字符
                const inputEvent = new InputEvent('input', { // 创建 input 事件
                    inputType: 'insertText',
                    data: char,
                    dataTransfer: null,
                    isComposing: false
                });
                inputElement.value += char; // 将当前字符添加到文本输入框中
                inputElement.dispatchEvent(inputEvent); // 分派 input 事件
                index++; // 增加当前输入字符的索引
                setTimeout(simulateInput, 100); // 100 毫秒后模拟下一个字符的输入
            }
        }
        simulateInput(); // 开始模拟输入
    }
</script>

</html>
```
