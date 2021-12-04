    var coeatable = $('#COEA')
    var ccstable = $('#CCS')
    var eductable = $('#EDUCTABLE')
    window.onload = setup();

    function setup(){
        COEA()
        CCS()
        EDUC()
    }
    // College of Engineering and Architecture Table
    function COEA(){
        console.log('coea')
        var $querycoea = "/admin/api/d/course/read/?/&course_type__icontains=Engineering"

        // For displaying Query
        $.get($querycoea, function (response)
        {
            var courses = response.result;
            $.each(courses, function(_, courses)
            {
                coeatable.append(
                    $("<tr>").append(
                        $("<li>").attr("class", "text-uppercase text-secondary text-m").text(courses.CourseDescription),
                    )
                )
            });
        }, "json");
    }
    // CCS TABLE
    function CCS(){
        console.log('ccs')
        var $queryccs = "/admin/api/d/course/read/?/&course_type__icontains=Computer%20Sciences"

        // For displaying Query
        $.get($queryccs, function (response)
        {
            var courses = response.result;
            $.each(courses, function(_, courses)
            {
                ccstable.append(
                    $("<tr>").append(
                        $("<li>").attr("class", "text-uppercase text-secondary text-m").text(courses.CourseDescription),
                    )
                )
            });
        }, "json");
    }

    // EDUC TABLE
    function EDUC(){
        console.log('ccs')
        var $queryeduc = "/admin/api/d/course/read/?/&course_type__icontains=Education"

        // For displaying Query
        $.get($queryeduc, function (response)
        {
            var courses = response.result;
            $.each(courses, function(_, courses)
            {
                eductable.append(
                    $("<tr>").append(
                        $("<li>").attr("class", "text-uppercase text-secondary text-m").text(courses.CourseDescription),
                    )
                )
            });
        }, "json");
    }

    // Add course Modal
    $("#add_course").on("click", function(){
        console.log("Click-Add Course");
        $("#addcourse_modal").fadeIn(100);
    })
    // Add Subject Modal
    $("#add_subject").on("click", function(){
        console.log("Click-Add Subject");
        $("#addsubject_modal").fadeIn(100);
    })

    // Add Course AJAX
    $("#addcourse_save").on("click", function(){
        $("#addcourse_modal").fadeOut(100);
        $.ajax({
            "method": "POST",
            "data": {
                "data-action": "new_course",
                "coursename": $("[name = 'coursename']").val(),
                "coursedesc": $("[name='coursedesc']").val(),
                "coursetype": $("[name='coursetype']").val(),
                "courseunits": $("[name='courseunits']").val(),
            },
            "success": function(data)
            {
                Swal.fire({
                    position: 'middle',
                    icon: 'success',
                    title: 'Course added succesfully',
                    showConfirmButton: false,
                    timer: 1500
                }).then((result) => {
                    $(location).attr('href', '/educ/courses/')
                });
            }
        })
    })