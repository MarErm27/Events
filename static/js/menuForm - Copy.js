(function(win, doc, $){
	"using strict";

	$(doc).ready(function(){
		var number = $("input[name='Number']");
		var cost = $("input[name='Cost']");
		var sum = $("input[name='Sum']");

		number.on("keyup", function(){
			sum.val(number.val() * cost.val());
		});

		cost.on("keyup", function(){
			sum.val(number.val() * cost.val());
		});
		var event = new Event('change');

		// Dispatch it.
		sum.dispatchEvent(event);
	});
})(window, document, $);