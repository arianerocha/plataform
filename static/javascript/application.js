// FormanticUI Particular Functions
$('.ui.checkbox').checkbox();

$('.message .close').on('click', function() {
    $(this).closest('.message').transition('fade')
});
