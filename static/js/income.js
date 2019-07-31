$(document).ready(function () {
    $("#name").on("blur focus", function () {
        if ($(this).is(":focus")) {
            $(".text").text("请下拉选择");

        } else {
            var name = $("#name").val();

            if (name.length === 0) {
                return
            }

            $.get("/income/query",
                {
                    name: name
                },
                function (data, status) {

                    // 清除多余选项
                    $("#cardDropdown").empty();

                    if (data.length > 2) {
                        $(".default.text").text("查出多个结果,点击选择");
                        $(".text").text("多个结果,点击选择");
                        $("#submit").attr("class", "ui secondary button");


                    } else if (data.length === 0) {
                        $(".default.text").text("未查询到");
                        $(".text").text("未查询到");
                        $("#submit").attr("class", "ui secondary disabled button");
                    }

                    // 循环增加选项至 cardDropdown div
                    for (var i = 0; i < data.length; i++) {
                        var div1 = "<div class=\"item\">" + data[i]["CardNumber"] + "</div>";

                        $("#cardDropdown").append(div1);
                    }

                });

        }
    });

});
// todo: 验证其余表单内容，是否为空