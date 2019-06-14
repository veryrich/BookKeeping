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
                                type: 'length[5]',
                                prompt: '账号必须5位以上'
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
                    },
                    confirmPassword: {
                        identifier: 'confirmPassword',
                        rules: [
                            {
                                type: 'match[password]',
                                prompt: '两次密码必须相同'
                            }
                        ]
                    }

                }
            });
        $('.ui.checkbox')
            .checkbox();

    });
