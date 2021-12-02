/* ===================================================================
 * Pudding - Main JS
 *
 * ------------------------------------------------------------------- */

(function($) {

    "use strict";
    
    var cfg = {
        scrollDuration : 800 // smoothscroll duration
    },

    $WIN = $(window);

    // Add the User Agent to the <html>
    // will be used for IE10 detection (Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0))
    var doc = document.documentElement;
    doc.setAttribute('data-useragent', navigator.userAgent);


    /* Preloader */
    var puddingPreloader = function() {
        
        $("html").addClass('cl-preload');

        $WIN.on('load', function() {

            //force page scroll position to top at page refresh
            // $('html, body').animate({ scrollTop: 0 }, 'normal');

            // will first fade out the loading animation 
            $("#loader").fadeOut("slow", function() {
                // will fade out the whole DIV that covers the website.
                $("#preloader").delay(300).fadeOut("slow");
            }); 
            
            // for hero content animations 
            $("html").removeClass('cl-preload');
            $("html").addClass('cl-loaded');
        
        });
    };


    /* Smooth Scrolling */
    var puddingSmoothScroll = function() {
        
        $('.smoothscroll').on('click', function (e) {
            var target = this.hash,
            $target    = $(target);
            
                e.preventDefault();
                e.stopPropagation();

            $('html, body').stop().animate({
                'scrollTop': $target.offset().top
            }, cfg.scrollDuration, 'swing').promise().done(function () {

                // check if menu is open
                if ($('body').hasClass('menu-is-open')) {
                    $('.header-menu-toggle').trigger('click');
                }

                window.location.hash = target;
            });
        });

    };


    /* Alert Boxes */
    var puddingAlertBoxes = function() {

        $('.alert-box').on('click', '.alert-box__close', function() {
            $(this).parent().fadeOut(500);
        }); 

    };


    /* Animate On Scroll */
    var puddingAOS = function() {
        
        AOS.init( {
            offset: 200,
            duration: 600,
            easing: 'ease-in-sine',
            delay: 300,
            once: true,
            disable: 'mobile'
        });

    };


    /* Back to Top */
    var puddingBackToTop = function() {
        
    var pxShow      = 500,
        goTopButton = $(".go-top")

        // Show or hide the button
        if ($(window).scrollTop() >= pxShow) goTopButton.addClass('link-is-visible');

        $(window).on('scroll', function() {
            if ($(window).scrollTop() >= pxShow) {
                if(!goTopButton.hasClass('link-is-visible')) goTopButton.addClass('link-is-visible')
            } else {
                goTopButton.removeClass('link-is-visible')
            }
        });

    };


    /* Tilt */
    var puddingTilt = function() {

        $('.pudding-tilt').tilt({
            glare: true,
            maxGlare: .25,
            perspective: 1000,
            maxTilt: 15
        });

    };

    /* Dropdown */
    var puddingDropdown = function() {

        document.addEventListener("click", e => {
            const isDropdownButton = e.target.matches("[data-dropdown-button]")
            if (!isDropdownButton && e.target.closest("[data-dropdown]") != null) return
          
            let currentDropdown
            if (isDropdownButton) {
              currentDropdown = e.target.closest("[data-dropdown]")
              currentDropdown.classList.toggle("active")
            }
          
            document.querySelectorAll("[data-dropdown].active").forEach(dropdown => {
              if (dropdown === currentDropdown) return
              dropdown.classList.remove("active")
            })
          })

    };

    /* Sidebar */
    var puddingSidebar = function() {
        var path = window.location.pathname;
        $('.sidebar-menu .sidebar-item.main-trigger').each(function ()
        {
            if ($(this).attr("href") == path)
            {
                $(this).addClass("active");
            }

            else
            {
                $(this).removeClass("active");
            }
        });

        $('.sidebar-menu .sidebar-item.main-trigger').unbind().on('click', function () {
            $('html, body').animate({ scrollTop: 0 }, 'fastest');
        })

        $('.sidebarBurger, .sidebar-close, .sidebar-action').unbind().on('click', function () {
            $('.page').toggleClass('toggle');
            $('.pudding-sidebar').toggleClass('active');
            $('.sidebar-action i').toggleClass('mdi-chevron-right mdi-chevron-left')
        })

        $('.openAccount').unbind().on('click', function() {
            $('.sidebar-dropdown').toggleClass('open');
            $('.sidebar-footer i.sidebar-collapsed-icon').toggleClass('mdi-unfold-more-horizontal mdi-unfold-less-horizontal');
        })
        
    };

    /* Test Modals */
    var puddingTestModals = function() {
        var modalOpen = function() {
            $('body').addClass('modal-open');
        };
        var noModalOpen = function() {
            $('body').removeClass('modal-open');
        };

        $('.openDefault').on('click', function() {
            $('#testdefault').addClass('modal-active');
            modalOpen();
        });

        $('.openDialog').on('click', function() {
            $('#testdialog').addClass('modal-active');
            modalOpen();
        });

        $('.openSmall').on('click', function() {
            $('#testsmall').addClass('modal-active');
            modalOpen();
        });

        $('.openMedium').on('click', function() {
            $('#testmedium').addClass('modal-active');
            modalOpen();
        });

        $('.openLarge').on('click', function() {
            $('#testlarge').addClass('modal-active');
            modalOpen();
        });

        $('.openFull').on('click', function() {
            $('#testfull').addClass('modal-active');
            modalOpen();
        });

        $('.openSheet').on('click', function() {
            $('#testbottomsheet').addClass('modal-active');
            modalOpen();
        });

        $('.openTop').on('click', function() {
            $('#testtop').addClass('modal-active');
            modalOpen();
        });

        $('.openRight').on('click', function() {
            $('#testright').addClass('modal-active');
            modalOpen();
        });

        $('.openBottom').on('click', function() {
            $('#testbottom').addClass('modal-active');
            modalOpen();
        });

        $('.openLeft').on('click', function() {
            $('#testleft').addClass('modal-active');
            modalOpen();
        });

        $('.close-modal').on('click', function() {
            $(this).closest('.modal').removeClass('modal-active');
            noModalOpen();
        });
    };

   /* Initialize
    * ------------------------------------------------------ */
    (function puddingInit() {

        puddingPreloader();
        puddingSmoothScroll();
        puddingAlertBoxes();
        puddingAOS();
        puddingBackToTop();
        puddingTilt();
        // puddingDropdown();
        puddingSidebar();
        puddingTestModals();
    })();

})(jQuery);