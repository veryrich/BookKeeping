$(document)
    .ready(function () {
        $('.ui.form')
            .form({
                fields: {
                    email: {
                        identifier: 'username',
                        rules: [
                            {
                                type: 'empty',
                                prompt: '用户名不能为空'
                            },
                            {
                                type: 'length[3]',
                                prompt: '账号必须3位以上'
                            }
                        ]
                    },
                    password: {
                        identifier: 'password',
                        rules: [
                            {
                                type: 'empty',
                                prompt: '密码不能为空'
                            },
                            {
                                type: 'length[5]',
                                prompt: '密码必须5位以上'
                            }
                        ]
                    }
                }
            });
    });