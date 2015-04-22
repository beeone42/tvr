

var tv_side = "left";


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
        $('#tv_left').click(function (){
            tv_side = "left";
            $(".side").html(tv_side);
            update_datas();
        });       
        $('#tv_right').click(function (){
            tv_side = "right";
            $(".side").html(tv_side);
            update_datas();
        });
        $('#publish').click(function (){
            publish_playlist(tv_side);
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

    //setInterval(function() { update_datas() }, 1000);

});

function update_datas()
{
    d = new Date();
    $('#slist').empty();
    $.getJSON("/ajax/state/" + tv_side, { t: d.getTime() }).done( function (data) {

        if (data.result.items && (data.result.items.length > 0))
        {
            $('#slist').empty();
            //alert(data.result.items[0].label);
            for (item of data.result.items)
            {
                $('#slist').append('<option>' + item.label + '</option>');
            }
        }
    });
}

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

function get_status()
{

}




