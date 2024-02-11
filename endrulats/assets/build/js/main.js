$(document).ready(function () {
    //store stuff
    var pane = $('#pane'),
        box = $('.box1'),
        boximg = $('.box1 img'),
        box2 = $('#box2'),
        box3 = $('#box3'),
        box4 = $('#box4'),
        wh = pane.width() - box.width(),
        wv = pane.height() - box.height(),
        d = {},
        x = 5;
    box.css("top", "20%");
    box.css("left", "20%");

    //ensure limits on the offsets
    function newh(v, a, b) {
        var n = parseInt(v, 10) - (d[a] ? x : 0) + (d[b] ? x : 0);
        return n < 0 ? 0 : n > wh ? wh : n;
    }

    function newv(v, a, b) {
        var n = parseInt(v, 10) - (d[a] ? x : 0) + (d[b] ? x : 0);
        return n < 0 ? 0 : n > wv ? wv : n;
    }

    //repaint
    setInterval(function () {
        box.css({
            left: function (i, v) {
                return newh(v, 37, 39);
            },
            top: function (i, v) {
                return newv(v, 38, 40);
            }
        });

        //sizes
        wh = pane.width() - box.width();
        wv = pane.height() - box.height();
    }, 90); //end of repaint

    //keys
    $(window).keydown(function (e) {
        switch (e.which) {
            case 39: // right
                boximg.attr('src', "img/game/img/walk.gif");
                break;
            case 38: // up
                boximg.attr('src', "img/game/img/up.gif");
                break;
            case 37: // left
                boximg.attr('src', "img/game/img/back.gif");
                break;
            case 40: // down
                boximg.attr('src', "img/game/img/down.gif");
                break;
            case 32: // spacebar
                var o = box;
                var bf = o.offset();
                var bfleft = bf.left;
                var bfright = bf.left + o.outerWidth();
                var bfbottom = bf.top + o.outerHeight();
                var bftop = bf.top;
                $(".d").each(function (index) {
                    var o2 = $(this);
                    var bf2 = o2.offset();
                    var bf2left = bf2.left;
                    var bf2right = bf2.left + o2.outerWidth();
                    var bf2bottom = bf2.top + o2.outerHeight();
                    var bf2top = bf2.top;
                    var box2 = $('#box2');

                    if (bfleft >= bf2left && bfright <= bf2right && bftop >= bf2top && bfbottom <= bf2bottom) {
                        if (o2.is(box2)) {
                            $(".zachhouse").addClass("active");
                            $(".zachhouse").addClass("active");
                            var messages = ['Welcome!', ' My name is Zachary Endrulat and I want to welcome you to my site.', 'We are going to take a trip to learning the best language ever! ']
                            $('#messages').chatBubble({
                                messages: messages,
                                typingSpeed: 40,
                                delay: 1000
                            });
                        }
                        if (o2.is(box4)) {
                            $(".beginpopup").addClass("active");
                            $(".popup-content").addClass("active");
                            var messages = ['Welcome!', 'You should begin your journey looking at these resources.', 'Please enjoy yourself!']
                            $('#messagesbegin').chatBubble({
                                messages: messages,
                                typingSpeed: 40,
                                delay: 1000,
                                
                            });
                            $('#messagesbegin').append('<center class="mescontainer"><a href="https://go.dev/tour/welcome/1" target="_blank"><button  class="fullbuttons btn-floating btn-large cyan pulse center close">Tour Of Go</button></a><br/>'+
                            '<a href="https://gobyexample.com" target="_blank"><button  class="fullbuttons btn-floating btn-large green pulse center close">Go By Example</button></a>'+
                            '<a href="https://go.dev/doc/effective_go" target="_blank"><button  class="fullbuttons btn-floating btn-large blue pulse center close">Effective Go</button></a>'+
                            '<br/><br/><center><button  class="fullbuttons btn-floating btn-large red pulse center close">Close</button></center>');

                        }

                        if (o2.is(box3)) {
                            $(".builderpopup").addClass("active");
                            $(".popup-content").addClass("active");
                            var messages = ['Welcome!', 'After you have gone over those resouces you should be ready to start building', 'Try to write a script that prints a struct.', 'remember <a href="https://gobyexample.com/structs" target="_blank">A struct is easy</a>']
                            $('#messagesbuilder').chatBubble({
                                messages: messages,
                                typingSpeed: 40,
                                delay: 1000,
                                
                            });
                            $('#messagesbuilder').append('<center >try to write the program in the <a href="https://go.dev/play/p/YfHt-B-NnIk" target="_blank">Playground</a></center> ');
                    }
                }
                });
                break;
            case 13: // entor

            default:
                boximg.attr('src', "/game/img/glogo.gif");
        }
        d[e.which] = true;
    });
    $(window).keyup(function (e) {
        d[e.which] = false;
    });
});


// has to be here for popup
popup();

bubble();