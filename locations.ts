export interface SubCountySeed {
    name: string;
    wards: string[];
    neighborhoods: string[];
}

export interface CountySeed {
    code: number;
    name: string;
    subCounties: SubCountySeed[];
    towns: string[];
}

export const KenyaLocationsData: CountySeed[] = [
    {
        code: 1, name: "Mombasa",
        towns: ["Mombasa City", "Nyali", "Bamburi", "Mvita", "Tudor", "Likoni", "Changamwe", "Shanzu"],
        subCounties: [
            { name: "Changamwe", wards: ["Changamwe", "Chaani", "Kipevu", "Airport", "Port Reitz"], neighborhoods: ["Chaani", "Kipevu", "Port Reitz"] },
            { name: "Jomvu", wards: ["Jomvu Kuu", "Miritini", "Mikindani"], neighborhoods: ["Mikindani", "Miritini"] },
            { name: "Kisauni", wards: ["MJamboni", "Junda", "Bamburi", "Mswambweni", "Magogoni"], neighborhoods: ["Bamburi", "Kisauni", "Shanzu"] },
            { name: "Nyali", wards: ["Frere Town", "Ziwani", "Mkomani", "Kongowea", "Kadzandani"], neighborhoods: ["Nyali Estate", "Kongowea", "Mkomani"] },
            { name: "Likoni", wards: ["Likoni", "Timbwani", "Bofu", "Bongwe", "Shika Adabu"], neighborhoods: ["Timbwani", "Bofu", "Shika Adabu"] },
            { name: "Mvita", wards: ["Mji wa Kale", "Tudor", "Tononoka", "Shimanzi", "Majengo"], neighborhoods: ["Tudor", "Tononoka", "Ganjoni"] }
        ]
    },
    {
        code: 2, name: "Kwale",
        towns: ["Ukunda", "Diani Beach", "Kwale", "Lunga Lunga", "Kinango", "Msambweni"],
        subCounties: [
            { name: "Msambweni", wards: ["Gombato", "Ukunda", "Kinondo", "Ramisi"], neighborhoods: ["Diani", "Ukunda Town"] },
            { name: "Lunga Lunga", wards: ["Pongwe", "Dzombo", "Mwereni", "Vanga"], neighborhoods: ["Lunga Lunga", "Vanga"] },
            { name: "Matuga", wards: ["Tsimba", "Waa", "Tiwi", "Kubo South"], neighborhoods: ["Tiwi", "Waa"] },
            { name: "Kinango", wards: ["Ndavaya", "Puma", "Kinango", "Mackinnon Road"], neighborhoods: ["Kinango Town", "Mackinnon"] }
        ]
    },
    {
        code: 3, name: "Kilifi",
        towns: ["Malindi", "Kilifi", "Watamu", "Vipingo", "Mtwapa", "Mariakani"],
        subCounties: [
            { name: "Kilifi North", wards: ["Tezo", "Sokoni", "Kibarani", "Dabaso", "Matsangoni"], neighborhoods: ["Kilifi Town", "Kibarani"] },
            { name: "Kilifi South", wards: ["Junju", "Mtwapa", "Chonyi", "Vipingo"], neighborhoods: ["Mtwapa Estate", "Vipingo Ridge"] },
            { name: "Malindi", wards: ["Jilore", "Kakuyuni", "Ganda", "Malindi Town", "Shella"], neighborhoods: ["Malindi Town", "Shella Estate"] },
            { name: "Magarini", wards: ["Marafa", "Magarini", "Gongoni", "Sabaki"], neighborhoods: ["Gongoni", "Sabaki"] },
            { name: "Kaloleni", wards: ["Kaloleni", "Kayafungo", "Mwanamwinga"], neighborhoods: ["Kaloleni Town"] },
            { name: "Rabai", wards: ["Mwawesa", "Ruruma", "Kambe", "Rabai"], neighborhoods: ["Rabai Town"] },
            { name: "Ganze", wards: ["Ganze", "Bamba", "Jaribuni", "Sokoke"], neighborhoods: ["Ganze Town", "Bamba"] }
        ]
    },
    {
        code: 4, name: "Tana River",
        towns: ["Hola", "Garsen", "Bura", "Madogo"],
        subCounties: [
            { name: "Garsen", wards: ["Kipini East", "Garsen South", "Kipini West"], neighborhoods: ["Garsen Town", "Kipini"] },
            { name: "Galole", wards: ["Wayu", "Chewani", "Mikinduni"], neighborhoods: ["Hola Town"] },
            { name: "Bura", wards: ["Chewele", "Bura", "Bangale"], neighborhoods: ["Bura Town", "Madogo"] }
        ]
    },
    {
        code: 5, name: "Lamu",
        towns: ["Lamu Town", "Shela", "Mpeketoni", "Mokowe", "Kiunga"],
        subCounties: [
            { name: "Lamu East", wards: ["Faza", "Kiunga", "Basuba"], neighborhoods: ["Faza", "Kiunga"] },
            { name: "Lamu West", wards: ["Shella", "Mkomani", "Hindi", "Mpeketoni", "Hongwe"], neighborhoods: ["Lamu Stone Town", "Shela Beach", "Mpeketoni"] }
        ]
    },
    {
        code: 6, name: "Taita/Taveta",
        towns: ["Voi", "Wundanyi", "Taveta", "Mwatate"],
        subCounties: [
            { name: "Taveta", wards: ["Chala", "Mahoo", "Bomani", "Mboghoni"], neighborhoods: ["Taveta Town"] },
            { name: "Wundanyi", wards: ["Wundanyi", "Werugha", "Mbale", "Wumingu"], neighborhoods: ["Wundanyi Town"] },
            { name: "Mwatate", wards: ["Rong'e", "Mwatate", "Bura", "Chavia"], neighborhoods: ["Mwatate Town"] },
            { name: "Voi", wards: ["Mbololo", "Sagalla", "Kaloleni", "Marungu"], neighborhoods: ["Voi Town", "Kaloleni Voi"] }
        ]
    },
    {
        code: 7, name: "Garissa",
        towns: ["Garissa", "Dadaab", "Masalani", "Bura East"],
        subCounties: [
            { name: "Garissa Township", wards: ["Waberi", "Iftin", "Township", "Galbet"], neighborhoods: ["Waberi", "Galbet", "Iftin"] },
            { name: "Balambala", wards: ["Balambala", "Sankuri", "Jarajara"], neighborhoods: ["Balambala Town"] },
            { name: "Lagdera", wards: ["Modogashe", "Benane", "Goreale"], neighborhoods: ["Modogashe"] },
            { name: "Dadaab", wards: ["Dadaab", "Labasigale", "Damajale"], neighborhoods: ["Dadaab Town"] },
            { name: "Fafi", wards: ["Bura", "Dekaharia", "Jarajila"], neighborhoods: ["Bura East"] },
            { name: "Ijara", wards: ["Ijara", "Masalani", "Hara"], neighborhoods: ["Masalani Town"] }
        ]
    },
    {
        code: 8, name: "Wajir",
        towns: ["Wajir", "Habaswein", "Bute", "Eldas"],
        subCounties: [
            { name: "Wajir North", wards: ["Bute", "Gurar", "Korondile"], neighborhoods: ["Bute Town"] },
            { name: "Wajir West", wards: ["Arbagajao", "Hadado", "Ademasajida"], neighborhoods: ["Hadado"] },
            { name: "Wajir East", wards: ["Wagberi", "Township", "Barmil"], neighborhoods: ["Wajir Town"] },
            { name: "Wajir South", wards: ["Benane", "Habaswein", "Lagboghol"], neighborhoods: ["Habaswein Town"] },
            { name: "Eldas", wards: ["Eldas", "Elben", "Dellow"], neighborhoods: ["Eldas Town"] },
            { name: "Tarbaj", wards: ["Tarbaj", "Elben", "Sarman"], neighborhoods: ["Tarbaj Town"] }
        ]
    },
    {
        code: 9, name: "Mandera",
        towns: ["Mandera", "Elwak", "Rhamu", "Takaba"],
        subCounties: [
            { name: "Mandera West", wards: ["Takaba", "Lagsure", "Gither"], neighborhoods: ["Takaba Town"] },
            { name: "Mandera Banissa", wards: ["Banissa", "Derkhale", "Guba"], neighborhoods: ["Banissa Town"] },
            { name: "Mandera North", wards: ["Rhamu", "Rhamu Dimtu", "Ashabito"], neighborhoods: ["Rhamu Town"] },
            { name: "Mandera South", wards: ["Elwak North", "Elwak South", "Shimbir Fatuma"], neighborhoods: ["Elwak Town"] },
            { name: "Mandera East", wards: ["Township", "Neboi", "Khalalio"], neighborhoods: ["Mandera Town"] },
            { name: "Lafey", wards: ["Lafey", "Waranqara", "Fino"], neighborhoods: ["Lafey Town"] }
        ]
    },
    {
        code: 10, name: "Marsabit",
        towns: ["Marsabit", "Moyale", "Laisamis", "Sololo", "North Horr"],
        subCounties: [
            { name: "Moyale", wards: ["Moyale Township", "Butiye", "Golbo", "Sololo"], neighborhoods: ["Moyale Town", "Butiye"] },
            { name: "North Horr", wards: ["Dukana", "North Horr", "Illeret"], neighborhoods: ["North Horr Town"] },
            { name: "Saku", wards: ["Sagante", "Karare", "Marsabit Central"], neighborhoods: ["Marsabit Town"] },
            { name: "Laisamis", wards: ["Laisamis", "Kargi", "Korr"], neighborhoods: ["Laisamis Town"] }
        ]
    },
    {
        code: 11, name: "Isiolo",
        towns: ["Isiolo", "Garbatulla", "Merti", "Oldonyiro"],
        subCounties: [
            { name: "Isiolo", wards: ["Wabera", "Bulla Pesa", "Burat", "Ngaremara"], neighborhoods: ["Wabera", "Bulla Pesa", "Isiolo Town"] },
            { name: "Garbatulla", wards: ["Garbatulla", "Kinna", "Sericho"], neighborhoods: ["Garbatulla Town", "Kinna"] },
            { name: "Merti", wards: ["Cherab", "Chari"], neighborhoods: ["Merti Town"] }
        ]
    },
    {
        code: 12, name: "Meru",
        towns: ["Meru", "Maua", "Timau", "Makutano", "Ntopic", "Laare"],
        subCounties: [
            { name: "Imenti South", wards: ["Mitunguu", "Igoji East", "Igoji West", "Abogeta East", "Abogeta West"], neighborhoods: ["Nkubu", "Igoji"] },
            { name: "Imenti North", wards: ["Municipality", "Ntima East", "Ntima West", "Nyaki East", "Nyaki West"], neighborhoods: ["Meru Town", "Makutano Meru"] },
            { name: "Imenti Central", wards: ["Mwanganthia", "Abothuguchi Central", "Abothuguchi West"], neighborhoods: ["Gaitu", "Central Imenti"] },
            { name: "Buuri", wards: ["Timau", "Kisima", "Ruiri/Rwarera"], neighborhoods: ["Timau Town", "Kisima"] },
            { name: "Tigania West", wards: ["Athwana", "Akithi", "Kianjai", "Nkomo"], neighborhoods: ["Kianjai", "Uringu"] },
            { name: "Tigania East", wards: ["Thangatha", "Mikinduri", "Kiguchwa"], neighborhoods: ["Mikinduri"] },
            { name: "Igembe South", wards: ["Maua", "Kiegoi", "Athiru Gaiti"], neighborhoods: ["Maua Town"] },
            { name: "Igembe Central", wards: ["Akirang'ondu", "Kanuni", "Njia"], neighborhoods: ["Laare"] },
            { name: "Igembe North", wards: ["Antuambui", "Ntunene", "Antubetwe Kiongo"], neighborhoods: ["Mutuati"] }
        ]
    },
    {
        code: 13, name: "Tharaka-Nithi",
        towns: ["Chuka", "Kathwana", "Chogoria", "Marimanti"],
        subCounties: [
            { name: "Chuka", wards: ["Mariani", "Karingani", "Magumoni"], neighborhoods: ["Chuka Town", "Karingani"] },
            { name: "Maara", wards: ["Muthambi", "Mwimbi", "Ganga"], neighborhoods: ["Chogoria Town"] },
            { name: "Tharaka", wards: ["Gatunga", "Mukothima", "Nkondi"], neighborhoods: ["Kathwana Town", "Marimanti"] }
        ]
    },
    {
        code: 14, name: "Embu",
        towns: ["Embu", "Runyenjes", "Siakago", "Kiritiri"],
        subCounties: [
            { name: "Manyatta", wards: ["Ruguru-Ngandori", "Kithimu", "Nginda", "Mbeti North"], neighborhoods: ["Embu Town", "Ngenge"] },
            { name: "Runyenjes", wards: ["Gaturi South", "Kagaari South", "Kagaari North", "Central Ward"], neighborhoods: ["Runyenjes Town"] },
            { name: "Mbeere North", wards: ["Nthawa", "Mevuriri", "Evurore"], neighborhoods: ["Siakago Town"] },
            { name: "Mbeere South", wards: ["Mwea", "Amantomba", "Mbeti South"], neighborhoods: ["Kiritiri Town"] }
        ]
    },
    {
        code: 15, name: "Kitui",
        towns: ["Kitui", "Mwingi", "Mutomo", "Kabati", "Kyuso"],
        subCounties: [
            { name: "Kitui Central", wards: ["Township", "Kyangwithya West", "Kyangwithya East", "Mulango"], neighborhoods: ["Kitui Town", "Mulango"] },
            { name: "Mwingi Central", wards: ["Central", "Kivou", "Nguni", "Nuu"], neighborhoods: ["Mwingi Town"] },
            { name: "Kitui South", wards: ["Ikutha", "Mutomo", "Kanziko"], neighborhoods: ["Mutomo Town"] },
            { name: "Kitui West", wards: ["Mutonguni", "Kauwi", "Matinyani"], neighborhoods: ["Kabati Town"] }
        ]
    },
    {
        code: 16, name: "Machakos",
        towns: ["Machakos", "Athi River", "Syokimau", "Mlolongo", "Kangundo", "Tala", "Matuu"],
        subCounties: [
            { name: "Machakos Town", wards: ["Kalama", "Muputi", "Machakos Central", "Muvuti/Kiima-Kimwe"], neighborhoods: ["Machakos Town", "Kiima Kimwe"] },
            { name: "Mavoko", wards: ["Athi River", "Syokimau/Mlolongo", "Muthwani", "Kinanie"], neighborhoods: ["Syokimau", "Mlolongo", "Athi River Estate"] },
            { name: "Kangundo", wards: ["Kangundo North", "Kangundo Central", "Kangundo East"], neighborhoods: ["Kangundo Town", "Tala"] },
            { name: "Yatta", wards: ["Ndalani", "Matuu", "Kithimani"], neighborhoods: ["Matuu Town"] },
            { name: "Mwala", wards: ["Mwala", "Makutano", "Masii"], neighborhoods: ["Masii", "Mwala"] }
        ]
    },
    {
        code: 17, name: "Makueni",
        towns: ["Wote", "Mtito Andei", "Kibwezi", "Emali", "Sultan Hamud"],
        subCounties: [
            { name: "Makueni", wards: ["Wote", "Muvau/Kikumini", "Mavindini"], neighborhoods: ["Wote Town"] },
            { name: "Kibwezi West", wards: ["Makindu", "Nguumo", "Emali/Mulala"], neighborhoods: ["Emali Town", "Makindu"] },
            { name: "Kibwezi East", wards: ["Masongaleni", "Mtito Andei", "Thange"], neighborhoods: ["Mtito Andei", "Kibwezi Town"] },
            { name: "Kilome", wards: ["Kasikeu", "Mukaa", "Kiima Kiu"], neighborhoods: ["Sultan Hamud"] }
        ]
    },
    {
        code: 18, name: "Nyandarua",
        towns: ["Ol Kalou", "Njabini", "Engineer", "Mairo Inya", "Ndaragua"],
        subCounties: [
            { name: "Ol Kalou", wards: ["Karau", "Kanjuiri Ridge", "Mirangine", "Kaimbaga"], neighborhoods: ["Ol Kalou Town"] },
            { name: "Kinangop", wards: ["Engineer", "Gathabai", "North Kinangop", "Njabini/Kiburu"], neighborhoods: ["Njabini", "Engineer Town"] },
            { name: "Ol Joro Orok", wards: ["Gathanji", "Gatimu", "Weru"], neighborhoods: ["Ol Joro Orok"] },
            { name: "Ndaragwa", wards: ["Leshau Pondo", "Kiriita", "Central"], neighborhoods: ["Mairo Inya"] }
        ]
    },
    {
        code: 19, name: "Nyeri",
        towns: ["Nyeri", "Karatina", "Othaya", "Mukurweini", "Naro Moru"],
        subCounties: [
            { name: "Nyeri Town", wards: ["Kiganjo/Mathari", "Rware", "Gatitu/Aguthi", "Ruring'u"], neighborhoods: ["Ruring'u", "Nyeri Town", "Ring Road"] },
            { name: "Mathira East", wards: ["Karatina Town", "Magutu", "Iria-ini"], neighborhoods: ["Karatina Town"] },
            { name: "Othaya", wards: ["Iria-ini", "Chinga", "Mahiga", "Karima"], neighborhoods: ["Othaya Town"] },
            { name: "Mukurweini", wards: ["Ruguru", "Gikondi", "Mukurweini West", "Mukurweini Central"], neighborhoods: ["Mukurweini Town"] },
            { name: "Kieni East", wards: ["Gakawa", "Naro Moru/Kiamathaga", "Thegu River"], neighborhoods: ["Naro Moru Town"] }
        ]
    },
    {
        code: 20, name: "Kirinyaga",
        towns: ["Kerugoya", "Sagana", "Kutus", "Wang'uru"],
        subCounties: [
            { name: "Kirinyaga Central", wards: ["Mutira", "Kaelo", "Kerugoya"], neighborhoods: ["Kerugoya Town"] },
            { name: "Mwea East", wards: ["Tebere", "Nyangati", "Murinduko"], neighborhoods: ["Wang'uru Town"] },
            { name: "Ndia", wards: ["Kariti", "Mukure", "Sagana"], neighborhoods: ["Sagana Town"] },
            { name: "Gichugu", wards: ["Kabare", "Baragwi", "Njukiini"], neighborhoods: ["Kutus Town"] }
        ]
    },
    {
        code: 21, name: "Murang'a",
        towns: ["Murang'a", "Kenol", "Maragua", "Kiria-ini", "Kangema"],
        subCounties: [
            { name: "Murang'a South", wards: ["Kimorori/Wempa", "Makuyu", "Kambiti"], neighborhoods: ["Kenol Town", "Makuyu"] },
            { name: "Maragua", wards: ["Ichagaki", "Nginda", "Makuyu"], neighborhoods: ["Maragua Town"] },
            { name: "Kiharu", wards: ["Township", "Mbiri", "Mugoiri"], neighborhoods: ["Murang'a Town"] },
            { name: "Kangema", wards: ["Kaniange", "Muguru", "Rwamuthambi"], neighborhoods: ["Kangema Town"] }
        ]
    },
    {
        code: 22, name: "Kiambu",
        towns: ["Thika", "Ruiru", "Kiambu", "Kikuyu", "Limuru", "Juja", "Karuri", "Githunguri"],
        subCounties: [
            { name: "Thika Town", wards: ["Township", "Kamenu", "Hospital", "Gatuanyaga"], neighborhoods: ["Makongeni", "Section 9", "Landless", "Happy Valley"] },
            { name: "Ruiru", wards: ["Gitothua", "Biashara", "Gatongora", "Kahawa Sukari", "Kahawa Wendani", "Kiuu"], neighborhoods: ["Kahawa Sukari", "Kahawa Wendani", "Mwihoko", "Membley"] },
            { name: "Juja", wards: ["Juja", "Kalimoni", "Witeithie", "Theta"], neighborhoods: ["Juja Town", "HighPoint", "Gachororo"] },
            { name: "Kiambu Town", wards: ["Township", "Ndumberi", "Riabai", "Ting'ang'a"], neighborhoods: ["Kiambu Town", "Ndumberi", "Indian Bazaar"] },
            { name: "Kabete", wards: ["Gitaru", "Muguga", "Nyadhuna", "Kabete", "Uthiru"], neighborhoods: ["Uthiru", "King'eero", "Lower Kabete"] },
            { name: "Kikuyu", wards: ["Karai", "Kikuyu", "Sigona", "Kinoo"], neighborhoods: ["Kinoo", "Muthiga", "Kikuyu Town"] },
            { name: "Limuru", wards: ["Limuru Central", "Ndeiya", "Limuru East", "Ngecha"], neighborhoods: ["Limuru Town", "Ngecha"] },
            { name: "Githunguri", wards: ["Githunguri", "Githiga", "Ikinu", "Ngewa"], neighborhoods: ["Githunguri Town", "Ikinu"] }
        ]
    },
    {
        code: 23, name: "Turkana",
        towns: ["Lodwar", "Kakuma", "Lokichogio", "Lokichar"],
        subCounties: [
            { name: "Turkana Central", wards: ["Kerio Delta", "Kang'atotha", "Lodwar Township"], neighborhoods: ["Lodwar Town"] },
            { name: "Turkana West", wards: ["Kakuma", "Lopur", "Letea"], neighborhoods: ["Kakuma Town"] },
            { name: "Turkana South", wards: ["Lokichar", "Katilu", "Kaputir"], neighborhoods: ["Lokichar Town"] }
        ]
    },
    {
        code: 24, name: "West Pokot",
        towns: ["Kapenguria", "Chepareria", "Makutano"],
        subCounties: [
            { name: "Kapenguria", wards: ["Riwo", "Kapenguria", "Endugh"], neighborhoods: ["Kapenguria Town", "Makutano"] },
            { name: "Pokot South", wards: ["Lelan", "Chepareria", "Batei"], neighborhoods: ["Chepareria Town"] }
        ]
    },
    {
        code: 25, name: "Samburu",
        towns: ["Maralal", "Baragoi", "Wamba", "Archers Post"],
        subCounties: [
            { name: "Samburu West", wards: ["Maralal", "Loosuk", "Poro"], neighborhoods: ["Maralal Town"] },
            { name: "Samburu East", wards: ["Wamba West", "Wamba East", "Waso"], neighborhoods: ["Wamba Town", "Archers Post"] }
        ]
    },
    {
        code: 26, name: "Trans Nzoia",
        towns: ["Kitale", "Kiminini", "Endebess"],
        subCounties: [
            { name: "Saboti", wards: ["Kitale Township", "Matisi", "Tuwani", "Saboti"], neighborhoods: ["Matisi", "Tuwani", "Milimani Kitale"] },
            { name: "Kiminini", wards: ["Kiminini", "Waitaluk", "Sirende"], neighborhoods: ["Kiminini Town"] },
            { name: "Endebess", wards: ["Endebess", "Matumbei", "Chepchoina"], neighborhoods: ["Endebess Town"] }
        ]
    },
    {
        code: 27, name: "Uasin Gishu",
        towns: ["Eldoret", "Turbo", "Burnt Forest", "Moiben"],
        subCounties: [
            { name: "Ainabkoi", wards: ["Kapsoya", "Kaptagat", "Ainabkoi"], neighborhoods: ["Kapsoya Estate", "Elgon View", "Kimumu"] },
            { name: "Kapseret", wards: ["Langas", "Simat/Kapseret", "Ngeria", "Megun"], neighborhoods: ["Langas Estate", "Pioneer", "Pioneer Estate"] },
            { name: "Turbo", wards: ["Ngenyilel", "Tapsagoi", "Kamagut", "Huruma"], neighborhoods: ["Huruma Estate", "Turbo Town"] },
            { name: "Moiben", wards: ["Tembelio", "Sergoit", "Karuna/Meibeki"], neighborhoods: ["Moiben Town"] }
        ]
    },
    {
        code: 28, name: "Elgeyo/Marakwet",
        towns: ["Iten", "Kapsowar", "Chebiemit"],
        subCounties: [
            { name: "Keiyo North", wards: ["Emsoo", "Kamariny", "Tambach"], neighborhoods: ["Iten Town", "Tambach"] },
            { name: "Marakwet West", wards: ["Kapsowar", "Lelan", "Sengwer"], neighborhoods: ["Kapsowar Town"] }
        ]
    },
    {
        code: 29, name: "Nandi",
        towns: ["Kapsabet", "Nandi Hills", "Mosoriot"],
        subCounties: [
            { name: "Emgwen", wards: ["Kapsabet", "Chepterwai", "Kilibwoni"], neighborhoods: ["Kapsabet Town"] },
            { name: "Nandi Hills", wards: ["Nandi Hills", "Chepkunyuk", "Ol'lessos"], neighborhoods: ["Nandi Hills Town"] },
            { name: "Chesumei", wards: ["Mosoriot", "Chemundu", "Kosirai"], neighborhoods: ["Mosoriot Town"] }
        ]
    },
    {
        code: 30, name: "Baringo",
        towns: ["Kabarnet", "Eldama Ravine", "Marigat"],
        subCounties: [
            { name: "Baringo Central", wards: ["Kabarnet", "Sacho", "Tenges"], neighborhoods: ["Kabarnet Town"] },
            { name: "Eldama Ravine", wards: ["Lembus", "Ravine", "Mumberes"], neighborhoods: ["Eldama Ravine Town"] },
            { name: "Baringo South", wards: ["Marigat", "Ilchamus", "Mochongoi"], neighborhoods: ["Marigat Town"] }
        ]
    },
    {
        code: 31, name: "Laikipia",
        towns: ["Nanyuki", "Nyahururu", "Rumuruti"],
        subCounties: [
            { name: "Laikipia East", wards: ["Nanyuki", "Umande", "Thingithu"], neighborhoods: ["Nanyuki Town", "Thingithu Estate"] },
            { name: "Laikipia West", wards: ["Nyahururu", "Rumuruti Township", "Githiga"], neighborhoods: ["Nyahururu Town", "Rumuruti"] }
        ]
    },
    {
        code: 32, name: "Nakuru",
        towns: ["Nakuru", "Naivasha", "Gilgil", "Molo", "Njoro"],
        subCounties: [
            { name: "Nakuru Town East", wards: ["Biashara", "Kivumbini", "Flamingo", "Menengai", "Nakuru East"], neighborhoods: ["Section 58", "Milimani Nakuru", "Free Area", "Lanet"] },
            { name: "Nakuru Town West", wards: ["Barut", "London", "Kaptembwo", "Kipkenyo", "Rhonda"], neighborhoods: ["Kaptembwo", "Rhonda", "London Estate"] },
            { name: "Naivasha", wards: ["Naivasha East", "Viwandani", "Mai Mahiu", "Maeilla", "Olkaria"], neighborhoods: ["Naivasha Town", "Mai Mahiu", "Karagita"] },
            { name: "Gilgil", wards: ["Gilgil", "Elementaita", "Mbaruk/Eburu"], neighborhoods: ["Gilgil Town"] },
            { name: "Molo", wards: ["Molo", "Mariashoni", "Elburgon"], neighborhoods: ["Molo Town", "Elburgon"] }
        ]
    },
    {
        code: 33, name: "Narok",
        towns: ["Narok", "Kilgoris"],
        subCounties: [
            { name: "Narok North", wards: ["Narok Town", "Nkareta", "Olokurto"], neighborhoods: ["Narok Town"] },
            { name: "Trans Mara West", wards: ["Kilgoris Central", "Keyian", "Angata Barikoi"], neighborhoods: ["Kilgoris Town"] }
        ]
    },
    {
        code: 34, name: "Kajiado",
        towns: ["Ngong", "Kitengela", "Ongata Rongai", "Kajiado", "Kiserian"],
        subCounties: [
            { name: "Kajiado North", wards: ["Ngong", "Ongata Rongai", "Nkaimurunya", "Oloolua"], neighborhoods: ["Ongata Rongai", "Ngong Town", "Kiserian", "Matasia"] },
            { name: "Kajiado East", wards: ["Kitengela", "Oloosirkon/Sholinke", "Kenyewa-Poka"], neighborhoods: ["Kitengela Town", "Acacia Estate"] },
            { name: "Kajiado Central", wards: ["Purko", "Ildamat", "Dalalekutuk"], neighborhoods: ["Kajiado Town"] }
        ]
    },
    {
        code: 35, name: "Kericho",
        towns: ["Kericho", "Kipkelion", "Litein"],
        subCounties: [
            { name: "Ainamoi", wards: ["Ainamoi", "Kapsoit", "Kipchebor"], neighborhoods: ["Kericho Town", "Kapsoit"] },
            { name: "Bureti", wards: ["Cheplanget", "Litein", "Cheboin"], neighborhoods: ["Litein Town"] }
        ]
    },
    {
        code: 36, name: "Bomet",
        towns: ["Bomet", "Sotik"],
        subCounties: [
            { name: "Bomet Central", wards: ["Siloam", "Ndaraweta", "Singorwet"], neighborhoods: ["Bomet Town"] },
            { name: "Sotik", wards: ["Ndanai/Abosi", "Chemagel", "Manaret/Rongena"], neighborhoods: ["Sotik Town"] }
        ]
    },
    {
        code: 37, name: "Kakamega",
        towns: ["Kakamega", "Mumias", "Butere"],
        subCounties: [
            { name: "Lurambi", wards: ["Butsotso East", "Butsotso South", "Sheywe"], neighborhoods: ["Kakamega Town", "Milimani Kakamega", "Amalemba"] },
            { name: "Mumias West", wards: ["Mumias Central", "Mumias North", "Etenje"], neighborhoods: ["Mumias Town"] },
            { name: "Butere", wards: ["Marama West", "Marama Central", "Marenyo"], neighborhoods: ["Butere Town"] }
        ]
    },
    {
        code: 38, name: "Vihiga",
        towns: ["Mbale", "Luanda", "Hamisi"],
        subCounties: [
            { name: "Vihiga", wards: ["Lugaga-Wamuluma", "South Maragoli"], neighborhoods: ["Mbale Town"] },
            { name: "Luanda", wards: ["Luanda Township", "Wemilabi", "Emabungo"], neighborhoods: ["Luanda Town"] }
        ]
    },
    {
        code: 39, name: "Bungoma",
        towns: ["Bungoma", "Webuye", "Kimilili"],
        subCounties: [
            { name: "Kanduyi", wards: ["Bungoma Township", "Bukembe West", "Khalaba"], neighborhoods: ["Bungoma Town", "Khalaba"] },
            { name: "Webuye West", wards: ["Sitikho", "Matulo", "Bokoli"], neighborhoods: ["Webuye Town"] },
            { name: "Kimilili", wards: ["Kimilili", "Kibingei", "Maeni"], neighborhoods: ["Kimilili Town"] }
        ]
    },
    {
        code: 40, name: "Busia",
        towns: ["Busia", "Malaba", "Nambale"],
        subCounties: [
            { name: "Matayos", wards: ["Busia Township", "Burumba", "Mayanje"], neighborhoods: ["Busia Town", "Burumba Estate"] },
            { name: "Teso North", wards: ["Malaba Central", "Malaba North", "Ang'urai"], neighborhoods: ["Malaba Town"] }
        ]
    },
    {
        code: 41, name: "Siaya",
        towns: ["Siaya", "Bondo", "Ugunja"],
        subCounties: [
            { name: "Alego Usonga", wards: ["Siaya Township", "Usonga", "North Alego"], neighborhoods: ["Siaya Town"] },
            { name: "Bondo", wards: ["Yimbo West", "Central Sakwa", "Bondo Township"], neighborhoods: ["Bondo Town"] },
            { name: "Ugunja", wards: ["Ugunja", "Sidindi", "Sigomere"], neighborhoods: ["Ugunja Town"] }
        ]
    },
    {
        code: 42, name: "Kisumu",
        towns: ["Kisumu City", "Muhoroni", "Ahero"],
        subCounties: [
            { name: "Kisumu Central", wards: ["Railways", "Migosi", "Shaurimoyo", "Kondele", "Market Milimani"], neighborhoods: ["Milimani Kisumu", "Migosi", "Kondele", "Riat Hills", "Mamboleo"] },
            { name: "Kisumu East", wards: ["Kajulu", "Kolwa East", "Manyatta B"], neighborhoods: ["Manyatta Estate", "Buoye"] },
            { name: "Nyando", wards: ["Ahero", "East Kano", "Awasi/Onjiko"], neighborhoods: ["Ahero Town"] }
        ]
    },
    {
        code: 43, name: "Homa Bay",
        towns: ["Homa Bay", "Mbita", "Oyugis"],
        subCounties: [
            { name: "Homa Bay Town", wards: ["Homa Bay Central", "Homa Bay Arujo", "Homa Bay West"], neighborhoods: ["Homa Bay Town", "Arujo Estate"] },
            { name: "Suba North", wards: ["Mbita", "Rusinga Island", "Gembe"], neighborhoods: ["Mbita Town"] },
            { name: "Kasipul", wards: ["Oyugis", "West Kasipul", "South Kasipul"], neighborhoods: ["Oyugis Town"] }
        ]
    },
    {
        code: 44, name: "Migori",
        towns: ["Migori", "Rongo", "Awendo", "Kehancha"],
        subCounties: [
            { name: "Suna West", wards: ["Wiga", "Wasweta II", "Ragana-Oruba"], neighborhoods: ["Migori Town", "Oruba Estate"] },
            { name: "Rongo", wards: ["Central Kamagambo", "East Kamagambo"], neighborhoods: ["Rongo Town"] },
            { name: "Awendo", wards: ["North Sakwa", "Central Sakwa"], neighborhoods: ["Awendo Town"] }
        ]
    },
    {
        code: 45, name: "Kisii",
        towns: ["Kisii", "Ogembo", "Suneka"],
        subCounties: [
            { name: "Nyaribari Chache", wards: ["Bobaracho", "Kisii Central", "Keumbu"], neighborhoods: ["Kisii Town", "Milimani Kisii", "Jogoo Estate"] },
            { name: "Bomachoge Chache", wards: ["Ogembo Town", "Majoge"], neighborhoods: ["Ogembo Town"] }
        ]
    },
    {
        code: 46, name: "Nyamira",
        towns: ["Nyamira", "Nyansiongo"],
        subCounties: [
            { name: "West Mugirango", wards: ["Nyamira Township", "Bogichora", "Bosamaro"], neighborhoods: ["Nyamira Town"] },
            { name: "Borabu", wards: ["Mekenene", "Nyansiongo", "Esise"], neighborhoods: ["Nyansiongo"] }
        ]
    },
    {
        code: 47, name: "Nairobi City",
        towns: ["Nairobi City"],
        subCounties: [
            { name: "Westlands", wards: ["Kitisuru", "Parklands/Highridge", "Karura", "Kangemi", "Mountain View"], neighborhoods: ["Westlands", "Parklands", "Spring Valley", "Brookside", "Kangemi"] },
            { name: "Dagoretti North", wards: ["Kilimani", "Kawangware", "Gatina", "Kileleshwa", "Kabiro"], neighborhoods: ["Kilimani", "Kileleshwa", "Lavington", "Kawangware"] },
            { name: "Dagoretti South", wards: ["Mutu-ini", "Ngando", "Riruta", "Uthiru/Ruthimitu", "Waithaka"], neighborhoods: ["Riruta", "Waithaka", "Dagoretti Corner"] },
            { name: "Lang'ata", wards: ["Karen", "Nairobi West", "Mugumo-ini", "South C", "Nyayo Highrise"], neighborhoods: ["Karen", "South C", "Nairobi West", "Dam Estate", "Langata"] },
            { name: "Kibra", wards: ["Laini Saba", "Lindi", "Makina", "Woodley/Kenyatta Golf Course", "Sarang'ombe"], neighborhoods: ["Woodley", "Golf Course Estate", "Adams Arcade"] },
            { name: "Roysambu", wards: ["Githurai", "Kahawa West", "Zimmerman", "Roysambu", "Kahawa"], neighborhoods: ["Roysambu", "Kahawa West", "Zimmerman", "TRM Area"] },
            { name: "Kasarani", wards: ["Clay City", "Mwiki", "Kasarani", "Njiru", "Ruai"], neighborhoods: ["Kasarani", "Mwiki", "Clay City", "Ruai", "Njiru"] },
            { name: "Ruaraka", wards: ["Babadogo", "Utalii", "Mathare North", "Lucky Summer", "Korogocho"], neighborhoods: ["Baba Dogo", "Lucky Summer", "Garden Estate"] },
            { name: "Embakasi South", wards: ["Imara Daima", "Kwa Njenga", "Kwa Reuben", "Pipeline", "Kware"], neighborhoods: ["Imara Daima", "Pipeline Estate"] },
            { name: "Embakasi North", wards: ["Kariobangi North", "Dandora Area I", "Dandora Area II", "Dandora Area III", "Dandora Area IV"], neighborhoods: ["Dandora", "Kariobangi"] },
            { name: "Embakasi Central", wards: ["Kayole North", "Kayole Central", "Kayole South", "Komarock", "Matopeni/Spring Valley"], neighborhoods: ["Komarock", "Kayole"] },
            { name: "Embakasi East", wards: ["Upper Savanna", "Lower Savanna", "Embakasi", "Utawala", "Mihango"], neighborhoods: ["Utawala", "Embakasi Village", "Mihango"] },
            { name: "Embakasi West", wards: ["Umoja I", "Umoja II", "Mowlem", "Donholm"], neighborhoods: ["Donholm", "Umoja", "Tena Estate", "Fedha"] },
            { name: "Makadara", wards: ["Maringo/Hamza", "Harambee", "Makadara", "Viwandani"], neighborhoods: ["Buruburu", "Makadara", "Hamza"] },
            { name: "Kamukunji", wards: ["Pumwani", "Eastleigh North", "Eastleigh South", "Airbase", "California"], neighborhoods: ["Eastleigh", "Pumwani"] },
            { name: "Starehe", wards: ["Nairobi Central", "Ngara", "Pangani", "Ziwani/Kariokor", "Landimawe", "Nairobi South"], neighborhoods: ["Nairobi CBD", "Ngara", "South B", "Pangani", "Park Road"] },
            { name: "Mathare", wards: ["Hospital", "Mabatini", "Huruma", "Ngei", "Mlango Kubwa", "Kiamaiko"], neighborhoods: ["Huruma", "Mathare"] }
        ]
    }
];
