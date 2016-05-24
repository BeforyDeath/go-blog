$("#pageForm").submit(function (e) {
    var form = $(this).serialize();

    var url = "/admin/page/create";
    var id = $(this).find("#id").val();
    if (id != 0) url = "/admin/page/update/" + id;

    $.ajax({
        type: "POST",
        url: url,
        data: form,
        success: function (data) {

            window.location.href = "/page/"+data;
            //console.log(data.alias)

        },
        error: function (data) {
            $(".error").text(data.responseText);
        }
    });
    e.preventDefault();
});