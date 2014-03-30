$("form").submit(function (event) {
  var email = $("input[name='email']").val();
  var password = $("input[name='password']").val();

  if (!email) {
    event.PreventDefault();
    alert("Please enter your email");
  }

  if (!password) {
    event.PreventDefault();
    alert("Please enter your password");
  }
});
