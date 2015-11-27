$(document).ready(function () {     	
	
	$("#resultslist").hide();
	
	/*---------------------------------------------------------------------
	|	      SEARCH_BUTTON              								   |	
	 ---------------------------------------------------------------------*/
	
	$("#btn_squery").on("click", function(){
		
		var _squery = $('#squery').val();		

		if(_squery !== "" && _squery !== "undefined"){
			
			var target = document.getElementById('spin_placeholder')
			var spinner = new Spinner(opts).spin(target);
			$("#spin_placeholder").show();
		
			seacrhQuery(_squery);
		}						
	});
	
	/*---------------------------------------------------------------------
	|		Sending SEARCH_QUERY to Server using AJAX POST REQUEST	       |	
	 ---------------------------------------------------------------------*/		
	
	function seacrhQuery(data){


		var sVal = { query : data };										
		
		$.ajax({
				type:"POST",
				url:"/requestSearchInfo",
				data: sVal,
				dataType:"json",
			success: function (data) {              
                console.log(data); 
				$('#serachlist').empty();							
				
				$('#serachlist').append( '<colgroup>'+
										 '<col class="col-md-1">'+
										 '<col class="col-md-5">' +	
										 '<col class="col-md-6">'+							
										 '</colgroup>'+
	                        			 '<th>Rank &nbsp;</th><th>Info</th>');

				$.each( data.Message.split('|'), function(index, value) {
				  console.log(value);
				  if ($.isNumeric(value)) {
					tRow = $('<tr style="cursor: pointer;">');
				  }
				  tCell = $('<td>').html(value);
				  $('#serachlist').append(tRow.append(tCell));
				});
				$("#spin_placeholder").hide();
            },
            error: function(xhr, textStatus, error){
		      console.log(xhr.statusText);
		      console.log(textStatus);
		      console.log(error);
 			}			
		});				
		clearInputVals("#squery");		
		$('#resultslist').show();
	}		

	/*---------------------------------------------------------------------
	|	 Pull first table-datacell value from Search Result List           |
	 ---------------------------------------------------------------------*/	
	
	$("#serachlist").on("click", "tr", function(e){		
		e.preventDefault();		
		var value = $(this).closest('tr').children('td:first').text();					 	
		console.log(value);
 	});	
	
	/*---------------------------------------------------------------------
	|		Clearing input values 										   |
	 ---------------------------------------------------------------------*/	
	function clearInputVals(tagElementID){			
		
		$('+ tagElementID +').val('');						
	}	
		
	/*---------------------------------------------------------------------
	|		Utility function for alert messages							   |		
	 ---------------------------------------------------------------------*/	
	function displayalert(message, alerttype){
		
		$("#alert_placeholder").append('<div id="alertdiv" class="alert ' + alerttype + '"><a class="close" data-dismiss="alert">×</a><span align="center">'+message+'</span></div>');
				
		window.setTimeout(function() {
		    $("#alertdiv").slideUp(1000, function(){
		        $(this).remove(); 
		    });
		}, 4000);
	}		
	
	/*---------------------------------------------------------------------
	|		Spinner for Progress indication    							   |
	 ---------------------------------------------------------------------*/
	
	var opts = {
				  lines: 10 // The number of lines to draw
				, length: 20 // The length of each line
				, width: 14 // The line thickness
				, radius: 26 // The radius of the inner circle
				, scale: 0.5 // Scales overall size of the spinner
				, corners: 0.8 // Corner roundness (0..1)
				, color: '#00CD00' // #rgb or #rrggbb or array of colors
				, opacity: 0.4 // Opacity of the lines
				, rotate: 0 // The rotation offset
				, direction: 1 // 1: clockwise, -1: counterclockwise
				, speed: 0.7 // Rounds per second
				, trail: 66 // Afterglow percentage
				, fps: 20 // Frames per second when using setTimeout() as a fallback for CSS
				, zIndex: 2e9 // The z-index (defaults to 2000000000)
				, className: 'spinner' // The CSS class to assign to the spinner
				, top: '51%' // Top position relative to parent
				, left: '50%' // Left position relative to parent
				, shadow: false // Whether to render a shadow
				, hwaccel: false // Whether to use hardware acceleration
				, position: 'relative' // Element positioning
				
				}	
	
});
