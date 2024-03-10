


$(document).ready(function () {
    document.addEventListener('DOMContentLoaded', function () {
        var elems = document.querySelectorAll('.sidenav');
        var instances = M.Sidenav.init(elems, options);
        var instance = M.Tabs.init(el, options);
    
    });

    
    $('.sidenav').sidenav();
    $('.tabs').tabs();
    var tooltipTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="tooltip"]'));
    var tooltipList = tooltipTriggerList.map(function (tooltipTriggerEl) {
        return new bootstrap.Tooltip(tooltipTriggerEl);
    });
  
    
});

