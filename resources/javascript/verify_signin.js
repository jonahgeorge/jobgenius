$("form").submit(function(event) {
event.preventDefault();

$.ajax({
type: 'POST',
url: '/api/signin',
data: { 
email    : $("input[name='email']").val(),
password : $("input[name='password']").val()
}
})
.done(function(data) {
location.href = '/';
})
.fail(function (data) {
if (data.responseText == "Unauthorized") {  
$(".alert").show().html("Invalid email or password.");
}
})
});

