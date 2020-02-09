var e = document.getElementById("Event")
menuTable = '<!-- Editable table -->\
<div class="card">\
  <h3 class="card-header text-center font-weight-bold text-uppercase py-4">Menu</h3>\
  <div class="card-body">\
    <div id="table" class="table-editable">\
      <span class="table-add float-right mb-3 mr-2"><a href="#!" class="text-success"><i\
            class="fas fa-plus fa-2x" aria-hidden="true"></i></a></span>\
      <table class="table table-bordered table-responsive-md table-striped text-center">\
        <thead>\
          <tr>\
            <th class="text-center">Nomenclature</th>\
            <th class="text-center">Number</th>\
            <th class="text-center">Cost</th>\
            <th class="text-center">Sum</th>\
            <th class="text-center">Kit</th>\
            <th class="text-center">Sort</th>\
            <th class="text-center">Remove</th>\
          </tr>\
        </thead>\
        <tbody>\
          <tr>\
            <td class="pt-3-half" <div class="container">\
            <div class="dropdown">\
              <button class="btn btn-primary dropdown-toggle" type="button" data-toggle="dropdown">Nomenclature\
              <span class="caret"></span></button>\
              <ul class="dropdown-menu">\
                <input class="form-control" id="myInput" type="text" placeholder="Search..">\
                <li><a href="#">Crisps</a></li>\
                <li><a href="#">Cola</a></li>\
                <li><a href="#">Burger</a></li>\
              </ul>\
            </div>\
          </div>\
          <script>\
          $(document).ready(function(){\
            $("#myInput").on("keyup", function() {\
              var value = $(this).val().toLowerCase();\
              $(".dropdown-menu li").filter(function() {\
                $(this).toggle($(this).text().toLowerCase().indexOf(value) > -1)\
              });\
            });\
          });\
          </script>\
            </td>\
            <td class="pt-3-half" contenteditable="true"></td>\
            <td class="pt-3-half" contenteditable="true"></td>\
            <td class="pt-3-half" contenteditable="true"></td>\
            <td class="pt-3-half" contenteditable="true"></td>\
            <td class="pt-3-half">\
              <span class="table-up"><a href="#!" class="indigo-text"><i class="fas fa-long-arrow-alt-up"\
                    aria-hidden="true"></i></a></span>\
              <span class="table-down"><a href="#!" class="indigo-text"><i class="fas fa-long-arrow-alt-down"\
                    aria-hidden="true"></i></a></span>\
            </td>\
            <td>\
              <span class="table-remove"><button type="button"\
                  class="btn btn-danger btn-rounded btn-sm my-0">Remove</button></span>\
            </td>\
          </tr>\
        </tbody>\
      </table>\
    </div>\
  </div>\
</div>\
<!-- Editable table -->'
e.innerHTML += menuTable
const $tableID = $('#table');
 const $BTN = $('#export-btn');
 const $EXPORT = $('#export');

 const newTr = `
<tr class="hide">
  <td class="pt-3-half" contenteditable="true">Example</td>
  <td class="pt-3-half" contenteditable="true">Example</td>
  <td class="pt-3-half" contenteditable="true">Example</td>
  <td class="pt-3-half" contenteditable="true">Example</td>
  <td class="pt-3-half" contenteditable="true">Example</td>
  <td class="pt-3-half">
    <span class="table-up"><a href="#!" class="indigo-text"><i class="fas fa-long-arrow-alt-up" aria-hidden="true"></i></a></span>
    <span class="table-down"><a href="#!" class="indigo-text"><i class="fas fa-long-arrow-alt-down" aria-hidden="true"></i></a></span>
  </td>
  <td>
    <span class="table-remove"><button type="button" class="btn btn-danger btn-rounded btn-sm my-0 waves-effect waves-light">Remove</button></span>
  </td>
</tr>`;

 $('.table-add').on('click', 'i', () => {

   const $clone = $tableID.find('tbody tr').last().clone(true).removeClass('hide table-line');

   if ($tableID.find('tbody tr').length === 0) {

     $('tbody').append(newTr);
   }

   $tableID.find('table').append($clone);
 });

 $tableID.on('click', '.table-remove', function () {

   $(this).parents('tr').detach();
 });

 $tableID.on('click', '.table-up', function () {

   const $row = $(this).parents('tr');

   if ($row.index() === 1) {
     return;
   }

   $row.prev().before($row.get(0));
 });

 $tableID.on('click', '.table-down', function () {

   const $row = $(this).parents('tr');
   $row.next().after($row.get(0));
 });

 // A few jQuery helpers for exporting only
 jQuery.fn.pop = [].pop;
 jQuery.fn.shift = [].shift;

 $BTN.on('click', () => {

   const $rows = $tableID.find('tr:not(:hidden)');
   const headers = [];
   const data = [];

   // Get the headers (add special header logic here)
   $($rows.shift()).find('th:not(:empty)').each(function () {

     headers.push($(this).text().toLowerCase());
   });

   // Turn all existing rows into a loopable array
   $rows.each(function () {
     const $td = $(this).find('td');
     const h = {};

     // Use the headers from earlier to name our hash keys
     headers.forEach((header, i) => {

       h[header] = $td.eq(i).text();
     });

     data.push(h);
   });

   // Output the result
   $EXPORT.text(JSON.stringify(data));
 });
 var dataFromAjax
 $.ajax({
    url: "http://localhost:8080/api/table",
    method: "GET",
    //data : { sendedData: 'Hello'},
     success : function(data) {
        alert(data);
        dataFromAjax = data;
        var result = JSON.parse(dataFromAjax);
        rov1 = $(document).find('#table').find('tr')[1].children;
        for (let i = 0; i < result.length; i++) {
          console.log(i);
          if (i == 0) {
            for (let j = 0; j < rov1.length; j++) {
              console.log(rov1[j].value);
            }
          }
        }
    },
  });