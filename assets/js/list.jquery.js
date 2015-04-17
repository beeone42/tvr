alert("lala");
$(function(){

	$.getJSON("/ajax/load/list").done(function(data){
		alert(data.id);
	});

});