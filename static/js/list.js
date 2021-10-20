$(document).ready(function () {
  var user_id
  let data={
    name :$.cookie('login')
  }
  //var user_id
  $.ajax({
    type: 'get',
    url: '/list/all',
    dataType: 'json',
    data: data,
    success: function (res) {
      //$(" #test").text(JSON.stringify(res))
      //alert(JSON.stringify(res));
      console.log(res.msg)
      console.log(res.id)
      user_id=res.id
      console.log($.cookie('login'))
      for (let i = 0; i < res.msg.length; i++) {
        if (res.msg[i].status == 0) {
          $('#todolist').append('<li class="list-group-item"> <div class="row">\
          <div class="col-md-10">' + res.msg[i].title + '</div>\
          <div class="col-md-1">\
            <button id="ok"  class="btn btn-success" style="border-radius: 50%;"><i class="bi bi-check"></i></button>\
          </div>\
          <div class="col-md-1">\
            <button id="delete" class="btn btn-danger" style="border-radius: 50%;"><i class="bi bi-x"></i></button>\
          </div></div></li>')
        }
        else{
          $('#todolist').append('<li class="list-group-item"> <div class="row">\
          <div class="col-md-10 text-muted" style="text-decoration:line-through">' + res.msg[i].title + '</div>\
          <div class="col-md-1">\
            <button id="notok" class="btn btn-warning" style="border-radius: 50%;"><i class="bi bi-check-all"></i></button>\
          </div>\
          <div class="col-md-1">\
            <button id="delete" class="btn btn-danger" style="border-radius: 50%;"><i class="bi bi-x"></i></button>\
          </div></div></li>')
        }
        //console.log($('li').eq(i).text())
        //$('li').eq(i).html(res[i].name)
      }
    }
  })

  $("#userdrop").html($.cookie('login'))

  $("#addlist").click(function(){
    var list ={
      id: 0,
      user_id: user_id, 
      title : $('#inputlist').val(),
      status : 0
    }
    console.log($('#inputlist').val())
    $.ajax({
      type: 'post',
      url: '/list',
      dataType: 'json',
      data: JSON.stringify(list),
      success: function (res) {
          alert(res.msg)
          window.location.reload()
      }
    })
  })

  $('#todolist').on('click','#ok,#notok',function(){
    str=$.trim($(this).parent().parent().text())
    let status;
    if ($(this).attr('id')=='ok')
      status= 0;
    else 
      status=1;
    let data ={
      title: str,
      status: status
    }
    $.ajax({
      type: 'put',
      url: '/list',
      dataType: 'json',
      data: data,
      success: function(res){
        alert(res.msg)
        window.location.reload()
      }
    })
  })
  $('#todolist').on('click','#delete',function(){
    str=$.trim($(this).parent().parent().text())
    let data ={
      title: str,
    }
    console.log(data)
    $.ajax({
      type: 'delete',
      url: '/list/'+str,
      dataType: 'json',
      data: data,
      success: function(res){
        alert(res.msg)
        window.location.reload()
      }
    })
  })
  $('#exit').click(function(){
    $.removeCookie('login',{ path: '/'});
    window.location.href = '/';
  })
})
