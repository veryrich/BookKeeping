$().ready(function () {

    // 删除按钮
    $(".ui.mini.red.button").click(function () {
        window.td_text = $(this).parent().prev().prev().prev().html();
        $('.ui.small.modal')
            .modal('show');
    });

    // 修改按钮
    $(".ui.mini.orange.button").click(function () {
        window.td_text = $(this).parent().prev().prev().prev().html();
        window.location.href = "/card/edit?cardNumber=" + window.td_text

    });

    $(".ui.positive.right.labeled.icon.button").click(function () {
        window.location.href = "/card/delete?cardNumber=" + window.td_text;
    });

    // 添加卡片表单验证
    $('#addCard')
        .form({
            fields: {
                email: {
                    identifier: 'name',
                    rules: [
                        {
                            type: 'empty',
                            prompt: '姓名不能为空'
                        }
                    ]
                },
                password: {
                    identifier: 'cardNumber',
                    rules: [
                        {
                            type: 'empty',
                            prompt: '卡号不能为空'
                        }
                    ]
                }
            }
        });


});

function goToCardLog() {
    window.location.assign("/card/log");
}

function goToAddCard() {
    window.location.assign("/card/add");
}
