(function (win, doc, $)
{
  // Set the Cookie by giving its cookie_name and cookie_value.
  function setCookie(cookie_name, cookie_value, cookie_expires)
  {
    document.cookie = cookie_name + "=" + cookie_value + ";expires=" + cookie_expires + "; path=/;SameSite=Lax";
  }

  // Getting the Cookie by declaring the cookie_name and will
  // return the cookie_value.
  function getCookie (cookie_name)
  {
    var document_cookie = document.cookie;
    var cookies = document_cookie.split(";");

    for (var i = 0, len = cookies.length; i < len; i++)
    {
      /* Checks if the cookie_name exists within
      * the list of cookies. If yes, then set the
      * cookie value to true.
      */
      var cookie = $.trim(cookies[i].indexOf(cookie_name)) != -1;

      // Return the cookie_value if found.
      if (cookie)
      {
        return cookie = cookies[i].split("=")[1];
      }
    }

    return "";
  }

  // Deleting the Cookie that equals to the cookie_name given
  // and setting the cookie_expires to know that the cookie is
  // already expired / will be deleted.
  function deleteCookie (cookie_name)
  {
    setCookie(cookie_name, "", new Date(+(new Date()) - (24 * 60 * 60 * 1000)).toUTCString());
  }

  function modalOpen()
  {
    $('body').addClass('modal-open');
  };

  function modalClose()
  {
    $('body').removeClass('modal-open');
  };

  function taskPriority(priority)
  {
    switch (priority)
    {
      case 1:
        return "Critical"
      case 2:
        return "Important"
      case 3:
        return "Normal"
      case 4:
        return "Low"
    }
  }

  function selectTaskPriority(select_element)
  {
    var options = [1, 2, 3, 4];    

    $.each(options, function (_, v)
    {
      select_element.append(
        $("<option>").attr("value", v)
                     .text(taskPriority(v))
      );
    });
  }

  function taskStatus(status)
  {
    switch (status)
    {
      case 1:
        return "Open"
      case 2:
        return "Pending"
      case 3:
        return "Closed"
      case 4:
        return "Rescheduled"
    }
  }

  function selectTaskStatus(select_element)
  {
    var options = [1, 2, 3, 4];    

    $.each(options, function (_, v)
    {
      select_element.append(
        $("<option>").attr("value", v)
                     .text(taskStatus(v))
      );
    });
  }

  function clearFields()
  {
    $("input,select,textarea").val("");
  }

  win["modal"] = {
    "open": modalOpen,
    "close": modalClose
  };

  win["common"] = {
    "cookie": {
      "set": setCookie,
      "get": getCookie,
      "delete": deleteCookie
    },
    "fields": {
      "clear": clearFields
    },
    "task": {
      "status": taskStatus,
      "priority": taskPriority,
      "select": {
        "status": selectTaskStatus,
        "priority": selectTaskPriority
      }
    }    
  };
})(window, document, $);