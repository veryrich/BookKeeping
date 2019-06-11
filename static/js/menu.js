$(document)
    .ready(function () {
        $('#dp1').dropdown();

        $(".item").click(function () {
            $(".active.item").removeClass("active");
            $(this).addClass("active");
        });

    });