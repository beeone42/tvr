

$(function(){

    d = new Date();

    $.getJSON("/ajax/video/", { t: d.getTime() }).done(function(data){
        //alert(data);
        for (tmp of data)
            $('#video').append('<option>' + tmp + '</option>');
    });

    $(document).on('change', "#plist", load_playlist);


    $.getJSON("/ajax/list/", { t: d.getTime() }).done(function(data){
        //alert(data);
        for (tmp of data)
            $('#plist').append('<option>' + tmp + '</option>');

        load_playlist();
    });

    if (navbar_active != "")
        $("#" + navbar_active).addClass("active");

    if (navbar_active == "main")
    {
        $('#publish_left').click(function (){
            publish_playlist("left");
        });       
        $('#publish_right').click(function (){
            publish_playlist("right");
        });
    }

    if (navbar_active == "navbar_video")
    {
        $('#playlist_add').click(function(event){
            event.preventDefault();
            var items = $("#video option:selected");
            var n = items.length;
            if (n > 0) {
                items.each(function(idx,item){
                    $("#lala").append('<option>' + item.text + '</option>');
                });
            }
            else {
                alert("Choose an item from list 1");
            }
        });
        $('#playlist_del').click(function(event){
            event.preventDefault();
            var items = $("#lala option:selected");
            var n = items.length;
            if (n > 0) {
                items.remove()
            }
            else {
                alert("Choose an item from list 1");
            }
        });
        $('#playlist_save').click(function(){
            $("#lala option").prop('selected', true);
        });
    }
});

function publish_playlist(tv)
{
    d = new Date();
    var items = $("#plist option:selected");

    if (items.length == 1)
    {
        //alert("publish playlist " + items[0].text + " on tv " + tv);
        $.get("/ajax/publish/" + tv + "/" + items[0].text, { t: d.getTime() });
    }
    else
    {
        alert("Select a playlist on the left first !");
    }
}


function load_playlist()
{
    d = new Date();
    var pname = $("#plist option:selected").text();
    if (pname == "")
        return ;

    $.getJSON("/ajax/load/" + pname, { t: d.getTime() }).done(function(data){
        $("#lala").empty();
        for (tmp of data.Items)
        {
            $('#lala').append('<option>' + tmp + '</option>');
        }
    });
}
