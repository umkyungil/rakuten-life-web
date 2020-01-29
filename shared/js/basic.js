var _is_pc;
var _$ele = {};

(function($){

    _$ele = {
        'document'      : $(document),
        'window'        : $(window),
        'body'          : $('body'),
        'header'        : $('#header'),
        'nav'           : $('#header').find('nav'),
        'spNavBtn'      : $('#header').find('.sp-nav-btn'),
        'pageTtl'       : $('#page-ttl'),
        'content'       : $('#content'),
        'footer'        : $('#footer'),
        'anchor'        : $('a[href^="#"]')
    };

    var scrollToPosition = function(id){
        var _opt = {
            'top'    : ($(id).offset() === null) ? 0 : $(id).offset().top,
            'speed'  : 600,
            'easing' : 'easeInOutQuart'
        };

        $('html, body')
            .stop(false, true)
            .animate({
                'scrollTop' : _opt.top
            }, {
                'duration' : _opt.speed,
                'easing'   : _opt.easing
            });
    };


    var toggleMenu = function(){
        var _opt = {
            'className' : 'open'
        };

        _$ele.nav
            .stop(false, true)
            .slideToggle();
        _$ele.spNavBtn
            .toggleClass(_opt.className);
    };

    $(function(){

        _$ele.window.setBreakpoints({
            distinct   : true,
            breakpoints: [1, 768]
        });

        _$ele.window.on('enterBreakpoint768', function(){
            _is_pc = true;
            _$ele.nav.removeAttr('style');
            _$ele.spNavBtn.removeClass('open');
        });

        _$ele.window.on('enterBreakpoint1', function(){
            _is_pc = false;
            _$ele.nav.removeAttr('style');
            _$ele.spNavBtn.removeClass('open');
        });

        _$ele.spNavBtn.on('click', function(){
            toggleMenu();
        });

        _$ele.anchor.on('click', function(e){
            e.preventDefault();
            scrollToPosition($(this).attr('href'));
        });


    });

})(jQuery);
