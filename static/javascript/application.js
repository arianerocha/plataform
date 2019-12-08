// FormanticUI Particular Functions
$('.ui.checkbox').checkbox();


// UX plataform Particular Functions

function plataformRemoveLoading() {
    $(".plataform.remove.loading").removeClass("loading");
}

setInterval(plataformRemoveLoading, 3000);
