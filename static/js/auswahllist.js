$(document).ready(function () {
	console.log("Auswahllist.js fired!")


	
	$("#anc_add").click(function(){
		console.log("function add fired!")
 	//	$('#tbl1 tr').last().after('<tr class="test"><td >Ich bin beschäftigt ['+cnt+']</td><tr>');
	
		var cnt = 1;	
		cnt++;
		var t=	$('#Tbldiv').append('<tr id="myrow"><td>Ich bin beschäftigt ['+cnt+']</td></tr>');
		console.log(t);
 	});
 
	$("#anc_rem").click(function(){
		console.log("function rem fired!")
		if($('#tbl1 tr').size()>1){
 		$('#tbl1 tr:last-child').remove();
		 }else{
 		alert('One row should be present in table');
		 }
  	});


	$('.Tbldiv').on("click",function() {
	$('#Tbldiv').addClass('myrow').siblings().removeClass('myrow');     //this.rowindex
   
	
		var x=document.getElementById('tbl1');
		var len=x.rows.length;
		console.log("rows --> " + len + " : " + x.rowIndex);
		
		
	
//	var rows = document.getElementById('tbl1').getElementsByTagName('tbody')[0].getElementsByTagName('tr');
	//var Xind;
	//for (i = 0; i < rows.length; i++) {
//        Xind = (this.rowIndex + 1);
//    }
//	alert("you have " + rows.length + " rows \n" + "Selected row index = " + Xind);
	});


});
	
	