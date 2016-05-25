$("#pageForm").submit(function (e) {
    var form = $(this).serialize();

    var url = "/admin/page/create";
    var id = $(this).find("#id").val();
    if (id != 0) url = "/admin/page/update";

    $.ajax({
        type: "POST",
        url: url,
        data: form,
        success: function (data) {
            window.location.href = "/page/" + data;
        },
        error: function (data) {
            $(".error").text(data.responseText);
        }
    });
    e.preventDefault();
});

function deletePage(id) {
    $.ajax({
        type: "POST",
        url: "/admin/page/delete",
        data: {id: id},
        success: function () {
            window.location.href = "/";
        },
        error: function (data) {
            console.log(data.responseText);
        }
    });
    return false
}

$("#description").keypress(function () {
    var text = $(this).val();
    $.ajax({
        type: "POST",
        url: "/admin/md",
        data: {text: text},
        success: function (data) {
            $("#md").html(data);
        },
        error: function (data) {
            console.log(data.responseText);
        }
    });
});