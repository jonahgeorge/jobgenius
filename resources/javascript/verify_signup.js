$("form").submit(function(event) {
  event.preventDefault();

  if ( $("input[name='password']").val().length < 8 || $("input[name='password']").val().length > 26 || $("input[name='password']").val().match(/\d+/g) == null) {
  $(".alert")
  .show()
  .text("Your password must be between 8 and 26 characters and have at least 1 number.");
  } else {
  if ( $("input[name='password']").val() == $("input[name='password_check']").val() )  {
    $.ajax({
      type : "POST",
      url  : "/api/signup",
      data : {
        email    : $("input[name='email']").val(),
        password : $("input[name='password']").val()
      }
    })
    .done(function (data) {
      if (data.status == "error") {
        $(".alert").show().text(data.message);
      } else if (data.status == "success") {
        window.location = "/";
      }
    })
    .fail(function (data) {
      $(".alert").show().text("An error occured. This isn't your fault, I'll have it fixed soon.");
    });
  } else { 
    $(".alert").show().text("Your passwords do not match.");
    $("input[name='password']").val("");
    $("input[name='password_check']").val("");
  }
  }
});

