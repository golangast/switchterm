

//collusion with a popup
 function Coll(o, o2, p) {
    var bf = o.offset();
    var bf2 = o2.offset();
    var bfleft = bf.left;
    var bfright = bf.left + o.outerWidth();
    var bfbottom = bf.top + o.outerHeight();
    var bftop = bf.top;
    var bf2left = bf2.left;
    var bf2right = bf2.left + o2.outerWidth();
    var bf2bottom = bf2.top + o2.outerHeight();
    var bf2top = bf2.top;
    console.log(bfleft, "<", bf2left, "&&", bfright, "<", bf2right, "&&", bftop, "<", bf2top, "&&", bfbottom, ">", bf2bottom)

    if (bfleft < bf2left && bfright < bf2right && bftop < bf2top && bfbottom > bf2bottom) {
        //crash
        $(p).addClass("active");
    } //end of collision

}

 function removeItemOnce(arr, value) {
    var index = arr.indexOf(value);
    if (index > -1) {
        arr.splice(index, 1);
    }
    return arr;
}

 function Release(o) {
    //put back on board
    pane.append(o);
    o.css("top", o.css('top'));
    o.css("left", o.css('left'));
    // if (guyarray.length > 0) {
    //     guyarray = removeItemOnce(guyarray, o);
    // } else {
    //     console.log(guyarray);
    // }
}

//collusion two obj
 function Collect(o, o2) {
    var bf = o.offset();
    var bf2 = o2.offset();
    var bfleft = bf.left;
    var bfright = bf.left + o.outerWidth();
    var bfbottom = bf.top + o.outerHeight();
    var bftop = bf.top;
    var bf2left = bf2.left;
    var bf2right = bf2.left + o2.outerWidth();
    var bf2bottom = bf2.top + o2.outerHeight();
    var bf2top = bf2.top;
    if (bfleft < bf2left && bfright < bf2right || bftop < bf2top && bfbottom > bf2bottom) {
        //crash
        o.append(o2);
        o2.css("top", 0);
        o2.css("left", 0);
    } //end of collision
}

 function Colls(o, o2) {
    var bf = o.offset();
    var bf2 = o2.offset();
    var bfleft = bf.left;
    var bfright = bf.left + o.outerWidth();
    var bfbottom = bf.top + o.outerHeight();
    var bftop = bf.top;
    var bf2left = bf2.left;
    var bf2right = bf2.left + o2.outerWidth();
    var bf2bottom = bf2.top + o2.outerHeight();
    var bf2top = bf2.top;
    console.log(bfleft, "<", bf2left, "&&", bfright, "<", bf2right, "&&", bftop, ">", bf2top, "&&", bfbottom, "<", bf2bottom)

    if (bfleft < bf2left && bfright < bf2right && bftop > bf2top && bfbottom < bf2bottom) {
        //crash

        if (o2.is(box2)) {
            $(".popup-overlay").addClass("active");
            $(".popup-content").addClass("active");
        }
    } //end of collision
}


 function popup(){

$(document).ready(function() {
    var text = {
        "question": [{
            "q": "What method do you use to print?",
            "a": ["Prnt", "p", "Println"],
            "r": "three"
        }, {
            "q": "sdfsd",
            "a": ["one", "two", "three"],
            "r": "two"
        }, {
            "q": "does sfdds work?",
            "a": ["one", "two", "three"],
            "r": "one"
        }, {
            "q": "does dsfdsffd work?",
            "a": ["one", "two", "three"],
            "r": "two"
        }]
    }

    //ramdomize
    function randomIntFromInterval(min, max) { // min and max included 
        return Math.floor(Math.random() * (max - min + 1) + min)
    }
    const rndInt = randomIntFromInterval(0, 3)

    $(".q").text(text.question[rndInt].q);
    $(".one").text(text.question[rndInt].a[0]);
    $(".two").text(text.question[rndInt].a[1]);
    $(".three").text(text.question[rndInt].a[2]);
    var r = text.question[rndInt].r

    $(".one").click(function() {
        if ("one" == r) {
            $(".one").before($(".g"));
            $(".g img").attr("src", "/static/img/hap.gif");
        } else {
            $(".g img").attr("src", "/static/img/mad.gif");
            $(".con").append($(".g"));
            $(".g ").css("width", "20%");
        }
    });
    $(".two").click(function() {
        if ("two" == r) {
            $(".two").before($(".g"));
            $(".g img").attr("src", "/static/img/hap.gif");

        } else {
            $(".g img").attr("src", "/static/img/mad.gif");
            $(".con").append($(".g"));
            $(".g ").css("width", "20%");
        }
    });
    $(".three").click(function() {
        if ("three" == r) {
            $(".three").before($(".g"));
            $(".g img").attr("src", "/static/img/hap.gif");

        } else {
            $(".g img").attr("src", "/static/img/mad.gif");
            $(".con").append($(".g"));
            $(".g ").css("width", "20%");
        }
    });

    $(".btn").click(function() {
        //ramdomize
        function randomIntFromInterval(min, max) { // min and max included 
            return Math.floor(Math.random() * (max - min + 1) + min)
        }
        const rndInt = randomIntFromInterval(0, 3)

        $(".q").text(text.question[rndInt].q);
        $(".one").text(text.question[rndInt].a[0]);
        $(".two").text(text.question[rndInt].a[1]);
        $(".three").text(text.question[rndInt].a[2]);
        var r = text.question[rndInt].r
        $(".con").append($(".g"));
        $(".g ").css("width", "20%");
        $(".g img").attr("src", "/static/img/gophers.png");

    });
});

}