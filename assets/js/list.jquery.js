$(function(){

    $.getJSON("/ajax/video").done(function(data){
        //alert(data);
        for (tmp of data)
            $('#video').append('<option>' + tmp + '</option>');
    });

    $(document).on('change', "#plist", load_playlist);


    $.getJSON("/ajax/list").done(function(data){
        //alert(data);
        for (tmp of data)
            $('#plist').append('<option>' + tmp + '</option>');

        load_playlist();
    });

});

function load_playlist()
{
    var pname = $("#plist option:selected").text();
    if (pname == "")
        return ;
    $.getJSON("/ajax/load/" + pname).done(function(data){
        $("#lala").empty();
        for (tmp of data.Items)
        {
            $('#lala').append('<option>' + tmp + '</option>');
        }
    });
}



