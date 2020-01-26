(function(win, doc, $){
	"using strict";

	$(doc).ready(function(){
		var number = $("input[name='Number']");
		var cost = $("input[name='Cost']");
		var sum = $("input[name='Sum']");
		var event = new Event('change');
		// Listen for the event.
		//sum[0].addEventListener('change', function (e) { alert("I've changed!") }, false);
		number.on("keyup", function(){
			sum.val(number.val() * cost.val());
		// Dispatch it.
		//sum[0].dispatchEvent(event);			
		});

		cost.on("keyup", function(){
			sum.val(number.val() * cost.val());
		// Dispatch it.
		//sum[0].dispatchEvent(event);			
		});


	});
})(window, document, $);