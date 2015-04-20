$(function(){

    $.getJSON("/ajax/load/video").done(function(data){
        //alert(data);
        for (tmp of data)
            $('#video').append('<option>' + tmp + '</option>');
    });

});
$(function(){

    $.getJSON("/ajax/load/list").done(function(data){
        //alert(data);
        for (tmp of data)
            $('#plist').append('<option>' + tmp + '</option>');
    });

});

$(document).on('change', "#plist", function(){
    var end = this.value;
 //  alert(end);
    $('#lala').empty().append('<option>' + end + '</option>');
});
