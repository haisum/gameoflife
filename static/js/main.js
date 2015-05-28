$(function(){
	var source   = $("#entry-template").html();
	var template = Handlebars.compile(source);
	var intervalSet = false;

	function refresh(){
		$.ajax({
			url : "/next",
			dataType: "json"
		}).done(function(data){
			$("#container").html(template(data));
			if(!intervalSet){
				window.setInterval(refresh, data.RefreshRate)
				intervalSet = true;
			}
		});
	}

	refresh();
});
