function alertType(type,message){
  console.log(type)
  $('#alert').empty()
  if (type==1){
    $('#alert').append('<div class="w-50 alert alert-primary mx-auto alert-dismissible fade show" role="alert">\
    <button type="button" class="btn-close" data-bs-dismiss="alert"></button><strong>'+message+'\
    </strong></div>')
  }else if(type==2){
    $('#alert').append('<div class="w-50 alert alert-info mx-auto alert-dismissible fade show" role="alert">\
    <button type="button" class="btn-close" data-bs-dismiss="alert"></button><strong>'+message+'\
    </strong></div>')
  }else if(type==3){
    $('#alert').append('<div class="w-50 alert alert-danger mx-auto alert-dismissible fade show" role="alert">\
    <button type="button" class="btn-close" data-bs-dismiss="alert"></button><strong>'+message+'\
    </strong></div>')
  }
}
$(document).ready(function () {
  $('#checkform').hide()
  $("#login").attr("disabled", true)
  $("#sign").attr("disabled", true)
  $('#signradio').click(function(){
    $("#login").attr("disabled", true)
    $("#sign").attr("disabled", false)
    $('#checkform').show()
    $('#floatingPassword').css({
      "margin-bottom": "-1px",
      "border-bottom-left-radius": "0",
      "border-bottom-right-radius": "0",
    })
  })
  $('#loginradio').click(function(){
    $("#sign").attr("disabled", true)
    $("#login").attr("disabled", false)
    $('#checkform').hide()
    $('#floatingPassword').css({
      "margin-bottom": "",
      "border-bottom-left-radius": "",
      "border-bottom-right-radius": "",
    })
  })
  $("#login").click(function () {
    console.log($('#form1').serialize())
    if($('#floatingInput').val()!=""&&$('#floatingPassword').val()!=""){
      $.ajax({
        type: 'post',
        url: '/login',
        dataType: 'json',
        data: $('#form1').serialize(),//序列化表单值
        success: function (res) {
            console.log(res)
            alertType(res.type,res.msg)
            if(res.type==1){
              setTimeout(function(){window.location.href = '/list'},900)
            }
        }
    })
    }else{
      alertType(2,"您有未输入的内容，请输入全部内容!")
    }
  })
  
  $('#sign').click(function(){
    let data = $('#form1').serializeArray()
    console.log(data)
    if($('#floatingInput').val()!=""&&$('#floatingPassword').val()!=""&&$('#floatingcheckPassword').val()!="" ){
      if(data[1].value!=data[2].value){
        alertType(2,"两次密码不一致!")
        $('#floatingcheckPassword').val('')
      }else if($('#floatingPassword').val().trim().length<8){
        alertType(2,"密码至少为8位，请重新输入!")
        $('#floatingcheckPassword').val('')
        $('#floatingPassword').val('')
      }else{
        $.ajax({
          type: 'post',
          url: '/sign',
          dataType: 'json',
          data: $('#form1').serialize(),//序列化表单值
          success: function (res) {
              console.log(res)
              alertType(res.type,res.msg)
          }
      })
      }
    }else{
      alertType(2,"您有未输入的内容，请输入全部内容!")
    }
  })
})