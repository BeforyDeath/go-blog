$("#page").submit(function (e) {
    var form = $(this).serialize();
    $.ajax({
        type: "POST",
        url: "/admin/page/create",
        data: form,
        success: function (data) {

            console.log(data);

        },
        error: function (data) {
            $(".error").text(data.responseText);
        }
    });
    e.preventDefault();
    return false
});