function page_view_load_video(data) {
    console.log(data);

    $('#page_view_content').show();
    $('#page_view_video_title').text(data.title);
    $('#page_view_video_channel').text(data.production.name).attr('href', '#/production/'+data.production.id);
}

function page_view(parameters) {
    console.log('ztj: page_view begin')

    video_id = parameters[0];
    console.log('ztj: query video info (vid:'+video_id+')')

    $.getJSON('/api/video/'+video_id).done(page_view_load_video);
}

function page_import() {
    console.debug('ztj: page_import begin');
    $('#page_import_content').show();
}

function main() {
    console.log('starting zobtube js')

    // check url
    url = $(location).attr('hash');

    // if url is too short, redirect to home
    if (url.length < 2) {
        window.location.replace("#/home")
    }

    url = url.substr(2); // remove '#/'
    parameters = url.split('/')
    target = parameters[0]
    parameters = parameters.slice(1)

    if (target == 'home') {
        return page_home(parameters);
    } else if (target == 'view') {
        return page_view(parameters);
    } else if (target == 'import') {
        return page_import(parameters);
    } else {
        // unknown, redirect to home
        window.location.replace("#/home")
    }
}

main()
