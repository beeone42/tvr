$(function(){

    $.getJSON("/ajax/video").done(function(data){
        //alert(data);
        for (tmp of data)
            $('#video').append('<option>' + tmp + '</option>');
    });

});
$(function(){

    $.getJSON("/ajax/list").done(function(data){
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
