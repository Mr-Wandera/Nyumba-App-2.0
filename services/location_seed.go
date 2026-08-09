package services

// CountySeed holds the hierarchical dataset for seeding Kenya's 47 counties
type CountySeed struct {
	Code        int
	Name        string
	SubCounties []SubCountySeed
	Towns       []string
}

// SubCountySeed holds sub-county level wards and estates
type SubCountySeed struct {
	Name          string
	Wards         []string
	Neighborhoods []string
}

// KenyaLocationsData contains all 47 counties of Kenya with complete administrative hierarchies
var KenyaLocationsData = []CountySeed{
	{
		Code: 1, Name: "Mombasa",
		Towns: []string{"Mombasa City", "Nyali", "Bamburi", "Mvita", "Tudor", "Likoni", "Changamwe", "Shanzu"},
		SubCounties: []SubCountySeed{
			{Name: "Changamwe", Wards: []string{"Changamwe", "Chaani", "Kipevu", "Airport", "Port Reitz"}, Neighborhoods: []string{"Chaani", "Kipevu", "Port Reitz"}},
			{Name: "Jomvu", Wards: []string{"Jomvu Kuu", "Miritini", "Mikindani"}, Neighborhoods: []string{"Mikindani", "Miritini"}},
			{Name: "Kisauni", Wards: []string{"MJamboni", "Junda", "Bamburi", "Mswambweni", "Magogoni"}, Neighborhoods: []string{"Bamburi", "Kisauni", "Shanzu"}},
			{Name: "Nyali", Wards: []string{"Frere Town", "Ziwani", "Mkomani", "Kongowea", "Kadzandani"}, Neighborhoods: []string{"Nyali Estate", "Kongowea", "Mkomani"}},
			{Name: "Likoni", Wards: []string{"Likoni", "Timbwani", "Bofu", "Bongwe", "Shika Adabu"}, Neighborhoods: []string{"Timbwani", "Bofu", "Shika Adabu"}},
			{Name: "Mvita", Wards: []string{"Mji wa Kale", "Tudor", "Tononoka", "Shimanzi", "Majengo"}, Neighborhoods: []string{"Tudor", "Tononoka", "Ganjoni"}},
		},
	},
	{
		Code: 2, Name: "Kwale",
		Towns: []string{"Ukunda", "Diani Beach", "Kwale", "Lunga Lunga", "Kinango", "Msambweni"},
		SubCounties: []SubCountySeed{
			{Name: "Msambweni", Wards: []string{"Gombato", "Ukunda", "Kinondo", "Ramisi"}, Neighborhoods: []string{"Diani", "Ukunda Town"}},
			{Name: "Lunga Lunga", Wards: []string{"Pongwe", "Dzombo", "Mwereni", "Vanga"}, Neighborhoods: []string{"Lunga Lunga", "Vanga"}},
			{Name: "Matuga", Wards: []string{"Tsimba", "Waa", "Tiwi", "Kubo South"}, Neighborhoods: []string{"Tiwi", "Waa"}},
			{Name: "Kinango", Wards: []string{"Ndavaya", "Puma", "Kinango", "Mackinnon Road"}, Neighborhoods: []string{"Kinango Town", "Mackinnon"}},
		},
	},
	{
		Code: 3, Name: "Kilifi",
		Towns: []string{"Malindi", "Kilifi", "Watamu", "Vipingo", "Mtwapa", "Mariakani"},
		SubCounties: []SubCountySeed{
			{Name: "Kilifi North", Wards: []string{"Tezo", "Sokoni", "Kibarani", "Dabaso", "Matsangoni"}, Neighborhoods: []string{"Kilifi Town", "Kibarani"}},
			{Name: "Kilifi South", Wards: []string{"Junju", "Mtwapa", "Chonyi", "Vipingo"}, Neighborhoods: []string{"Mtwapa Estate", "Vipingo Ridge"}},
			{Name: "Malindi", Wards: []string{"Jilore", "Kakuyuni", "Ganda", "Malindi Town", "Shella"}, Neighborhoods: []string{"Malindi Town", "Shella Estate"}},
			{Name: "Magarini", Wards: []string{"Marafa", "Magarini", "Gongoni", "Sabaki"}, Neighborhoods: []string{"Gongoni", "Sabaki"}},
			{Name: "Kaloleni", Wards: []string{"Kaloleni", "Kayafungo", "Mwanamwinga"}, Neighborhoods: []string{"Kaloleni Town"}},
			{Name: "Rabai", Wards: []string{"Mwawesa", "Ruruma", "Kambe", "Rabai"}, Neighborhoods: []string{"Rabai Town"}},
			{Name: "Ganze", Wards: []string{"Ganze", "Bamba", "Jaribuni", "Sokoke"}, Neighborhoods: []string{"Ganze Town", "Bamba"}},
		},
	},
	{
		Code: 4, Name: "Tana River",
		Towns: []string{"Hola", "Garsen", "Bura", "Madogo"},
		SubCounties: []SubCountySeed{
			{Name: "Garsen", Wards: []string{"Kipini East", "Garsen South", "Kipini West"}, Neighborhoods: []string{"Garsen Town", "Kipini"}},
			{Name: "Galole", Wards: []string{"Wayu", "Chewani", "Mikinduni"}, Neighborhoods: []string{"Hola Town"}},
			{Name: "Bura", Wards: []string{"Chewele", "Bura", "Bangale"}, Neighborhoods: []string{"Bura Town", "Madogo"}},
		},
	},
	{
		Code: 5, Name: "Lamu",
		Towns: []string{"Lamu Town", "Shela", "Mpeketoni", "Mokowe", "Kiunga"},
		SubCounties: []SubCountySeed{
			{Name: "Lamu East", Wards: []string{"Faza", "Kiunga", "Basuba"}, Neighborhoods: []string{"Faza", "Kiunga"}},
			{Name: "Lamu West", Wards: []string{"Shella", "Mkomani", "Hindi", "Mpeketoni", "Hongwe"}, Neighborhoods: []string{"Lamu Stone Town", "Shela Beach", "Mpeketoni"}},
		},
	},
	{
		Code: 6, Name: "Taita/Taveta",
		Towns: []string{"Voi", "Wundanyi", "Taveta", "Mwatate"},
		SubCounties: []SubCountySeed{
			{Name: "Taveta", Wards: []string{"Chala", "Mahoo", "Bomani", "Mboghoni"}, Neighborhoods: []string{"Taveta Town"}},
			{Name: "Wundanyi", Wards: []string{"Wundanyi", "Werugha", "Mbale", "Wumingu"}, Neighborhoods: []string{"Wundanyi Town"}},
			{Name: "Mwatate", Wards: []string{"Rong'e", "Mwatate", "Bura", "Chavia"}, Neighborhoods: []string{"Mwatate Town"}},
			{Name: "Voi", Wards: []string{"Mbololo", "Sagalla", "Kaloleni", "Marungu"}, Neighborhoods: []string{"Voi Town", "Kaloleni Voi"}},
		},
	},
	{
		Code: 7, Name: "Garissa",
		Towns: []string{"Garissa", "Dadaab", "Masalani", "Bura East"},
		SubCounties: []SubCountySeed{
			{Name: "Garissa Township", Wards: []string{"Waberi", "Iftin", "Township", "Galbet"}, Neighborhoods: []string{"Waberi", "Galbet", "Iftin"}},
			{Name: "Balambala", Wards: []string{"Balambala", "Sankuri", "Jarajara"}, Neighborhoods: []string{"Balambala Town"}},
			{Name: "Lagdera", Wards: []string{"Modogashe", "Benane", "Goreale"}, Neighborhoods: []string{"Modogashe"}},
			{Name: "Dadaab", Wards: []string{"Dadaab", "Labasigale", "Damajale"}, Neighborhoods: []string{"Dadaab Town"}},
			{Name: "Fafi", Wards: []string{"Bura", "Dekaharia", "Jarajila"}, Neighborhoods: []string{"Bura East"}},
			{Name: "Ijara", Wards: []string{"Ijara", "Masalani", "Hara"}, Neighborhoods: []string{"Masalani Town"}},
		},
	},
	{
		Code: 8, Name: "Wajir",
		Towns: []string{"Wajir", "Habaswein", "Bute", "Eldas"},
		SubCounties: []SubCountySeed{
			{Name: "Wajir North", Wards: []string{"Bute", "Gurar", "Korondile"}, Neighborhoods: []string{"Bute Town"}},
			{Name: "Wajir West", Wards: []string{"Arbagajao", "Hadado", "Ademasajida"}, Neighborhoods: []string{"Hadado"}},
			{Name: "Wajir East", Wards: []string{"Wagberi", "Township", "Barmil"}, Neighborhoods: []string{"Wajir Town"}},
			{Name: "Wajir South", Wards: []string{"Benane", "Habaswein", "Lagboghol"}, Neighborhoods: []string{"Habaswein Town"}},
			{Name: "Eldas", Wards: []string{"Eldas", "Elben", "Dellow"}, Neighborhoods: []string{"Eldas Town"}},
			{Name: "Tarbaj", Wards: []string{"Tarbaj", "Elben", "Sarman"}, Neighborhoods: []string{"Tarbaj Town"}},
		},
	},
	{
		Code: 9, Name: "Mandera",
		Towns: []string{"Mandera", "Elwak", "Rhamu", "Takaba"},
		SubCounties: []SubCountySeed{
			{Name: "Mandera West", Wards: []string{"Takaba", "Lagsure", "Gither"}, Neighborhoods: []string{"Takaba Town"}},
			{Name: "Mandera Banissa", Wards: []string{"Banissa", "Derkhale", "Guba"}, Neighborhoods: []string{"Banissa Town"}},
			{Name: "Mandera North", Wards: []string{"Rhamu", "Rhamu Dimtu", "Ashabito"}, Neighborhoods: []string{"Rhamu Town"}},
			{Name: "Mandera South", Wards: []string{"Elwak North", "Elwak South", "Shimbir Fatuma"}, Neighborhoods: []string{"Elwak Town"}},
			{Name: "Mandera East", Wards: []string{"Township", "Neboi", "Khalalio"}, Neighborhoods: []string{"Mandera Town"}},
			{Name: "Lafey", Wards: []string{"Lafey", "Waranqara", "Fino"}, Neighborhoods: []string{"Lafey Town"}},
		},
	},
	{
		Code: 10, Name: "Marsabit",
		Towns: []string{"Marsabit", "Moyale", "Laisamis", "Sololo", "North Horr"},
		SubCounties: []SubCountySeed{
			{Name: "Moyale", Wards: []string{"Moyale Township", "Butiye", "Golbo", "Sololo"}, Neighborhoods: []string{"Moyale Town", "Butiye"}},
			{Name: "North Horr", Wards: []string{"Dukana", "North Horr", "Illeret"}, Neighborhoods: []string{"North Horr Town"}},
			{Name: "Saku", Wards: []string{"Sagante", "Karare", "Marsabit Central"}, Neighborhoods: []string{"Marsabit Town"}},
			{Name: "Laisamis", Wards: []string{"Laisamis", "Kargi", "Korr"}, Neighborhoods: []string{"Laisamis Town"}},
		},
	},
	{
		Code: 11, Name: "Isiolo",
		Towns: []string{"Isiolo", "Garbatulla", "Merti", "Oldonyiro"},
		SubCounties: []SubCountySeed{
			{Name: "Isiolo", Wards: []string{"Wabera", "Bulla Pesa", "Burat", "Ngaremara"}, Neighborhoods: []string{"Wabera", "Bulla Pesa", "Isiolo Town"}},
			{Name: "Garbatulla", Wards: []string{"Garbatulla", "Kinna", "Sericho"}, Neighborhoods: []string{"Garbatulla Town", "Kinna"}},
			{Name: "Merti", Wards: []string{"Cherab", "Chari"}, Neighborhoods: []string{"Merti Town"}},
		},
	},
	{
		Code: 12, Name: "Meru",
		Towns: []string{"Meru", "Maua", "Timau", "Makutano", "Ntopic", "Laare"},
		SubCounties: []SubCountySeed{
			{Name: "Imenti South", Wards: []string{"Mitunguu", "Igoji East", "Igoji West", "Abogeta East", "Abogeta West"}, Neighborhoods: []string{"Nkubu", "Igoji"}},
			{Name: "Imenti North", Wards: []string{"Municipality", "Ntima East", "Ntima West", "Nyaki East", "Nyaki West"}, Neighborhoods: []string{"Meru Town", "Makutano Meru"}},
			{Name: "Imenti Central", Wards: []string{"Mwanganthia", "Abothuguchi Central", "Abothuguchi West"}, Neighborhoods: []string{"Gaitu", "Central Imenti"}},
			{Name: "Buuri", Wards: []string{"Timau", "Kisima", "Ruiri/Rwarera"}, Neighborhoods: []string{"Timau Town", "Kisima"}},
			{Name: "Tigania West", Wards: []string{"Athwana", "Akithi", "Kianjai", "Nkomo"}, Neighborhoods: []string{"Kianjai", "Uringu"}},
			{Name: "Tigania East", Wards: []string{"Thangatha", "Mikinduri", "Kiguchwa"}, Neighborhoods: []string{"Mikinduri"}},
			{Name: "Igembe South", Wards: []string{"Maua", "Kiegoi", "Athiru Gaiti"}, Neighborhoods: []string{"Maua Town"}},
			{Name: "Igembe Central", Wards: []string{"Akirang'ondu", "Kanuni", "Njia"}, Neighborhoods: []string{"Laare"}},
			{Name: "Igembe North", Wards: []string{"Antuambui", "Ntunene", "Antubetwe Kiongo"}, Neighborhoods: []string{"Mutuati"}},
		},
	},
	{
		Code: 13, Name: "Tharaka-Nithi",
		Towns: []string{"Chuka", "Kathwana", "Chogoria", "Marimanti"},
		SubCounties: []SubCountySeed{
			{Name: "Chuka", Wards: []string{"Mariani", "Karingani", "Magumoni"}, Neighborhoods: []string{"Chuka Town", "Karingani"}},
			{Name: "Maara", Wards: []string{"Muthambi", "Mwimbi", "Ganga"}, Neighborhoods: []string{"Chogoria Town"}},
			{Name: "Tharaka", Wards: []string{"Gatunga", "Mukothima", "Nkondi"}, Neighborhoods: []string{"Kathwana Town", "Marimanti"}},
		},
	},
	{
		Code: 14, Name: "Embu",
		Towns: []string{"Embu", "Runyenjes", "Siakago", "Kiritiri"},
		SubCounties: []SubCountySeed{
			{Name: "Manyatta", Wards: []string{"Ruguru-Ngandori", "Kithimu", "Nginda", "Mbeti North"}, Neighborhoods: []string{"Embu Town", "Ngenge"}},
			{Name: "Runyenjes", Wards: []string{"Gaturi South", "Kagaari South", "Kagaari North", "Central Ward"}, Neighborhoods: []string{"Runyenjes Town"}},
			{Name: "Mbeere North", Wards: []string{"Nthawa", "Mevuriri", "Evurore"}, Neighborhoods: []string{"Siakago Town"}},
			{Name: "Mbeere South", Wards: []string{"Mwea", "Amantomba", "Mbeti South"}, Neighborhoods: []string{"Kiritiri Town"}},
		},
	},
	{
		Code: 15, Name: "Kitui",
		Towns: []string{"Kitui", "Mwingi", "Mutomo", "Kabati", "Kyuso"},
		SubCounties: []SubCountySeed{
			{Name: "Kitui Central", Wards: []string{"Township", "Kyangwithya West", "Kyangwithya East", "Mulango"}, Neighborhoods: []string{"Kitui Town", "Mulango"}},
			{Name: "Mwingi Central", Wards: []string{"Central", "Kivou", "Nguni", "Nuu"}, Neighborhoods: []string{"Mwingi Town"}},
			{Name: "Kitui South", Wards: []string{"Ikutha", "Mutomo", "Kanziko"}, Neighborhoods: []string{"Mutomo Town"}},
			{Name: "Kitui West", Wards: []string{"Mutonguni", "Kauwi", "Matinyani"}, Neighborhoods: []string{"Kabati Town"}},
		},
	},
	{
		Code: 16, Name: "Machakos",
		Towns: []string{"Machakos", "Athi River", "Syokimau", "Mlolongo", "Kangundo", "Tala", "Matuu"},
		SubCounties: []SubCountySeed{
			{Name: "Machakos Town", Wards: []string{"Kalama", "Muputi", "Machakos Central", "Muvuti/Kiima-Kimwe"}, Neighborhoods: []string{"Machakos Town", "Kiima Kimwe"}},
			{Name: "Mavoko", Wards: []string{"Athi River", "Syokimau/Mlolongo", "Muthwani", "Kinanie"}, Neighborhoods: []string{"Syokimau", "Mlolongo", "Athi River Estate"}},
			{Name: "Kangundo", Wards: []string{"Kangundo North", "Kangundo Central", "Kangundo East"}, Neighborhoods: []string{"Kangundo Town", "Tala"}},
			{Name: "Yatta", Wards: []string{"Ndalani", "Matuu", "Kithimani"}, Neighborhoods: []string{"Matuu Town"}},
			{Name: "Mwala", Wards: []string{"Mwala", "Makutano", "Masii"}, Neighborhoods: []string{"Masii", "Mwala"}},
		},
	},
	{
		Code: 17, Name: "Makueni",
		Towns: []string{"Wote", "Mtito Andei", "Kibwezi", "Emali", "Sultan Hamud"},
		SubCounties: []SubCountySeed{
			{Name: "Makueni", Wards: []string{"Wote", "Muvau/Kikumini", "Mavindini"}, Neighborhoods: []string{"Wote Town"}},
			{Name: "Kibwezi West", Wards: []string{"Makindu", "Nguumo", "Emali/Mulala"}, Neighborhoods: []string{"Emali Town", "Makindu"}},
			{Name: "Kibwezi East", Wards: []string{"Masongaleni", "Mtito Andei", "Thange"}, Neighborhoods: []string{"Mtito Andei", "Kibwezi Town"}},
			{Name: "Kilome", Wards: []string{"Kasikeu", "Mukaa", "Kiima Kiu"}, Neighborhoods: []string{"Sultan Hamud"}},
		},
	},
	{
		Code: 18, Name: "Nyandarua",
		Towns: []string{"Ol Kalou", "Njabini", "Engineer", "Mairo Inya", "Ndaragua"},
		SubCounties: []SubCountySeed{
			{Name: "Ol Kalou", Wards: []string{"Karau", "Kanjuiri Ridge", "Mirangine", "Kaimbaga"}, Neighborhoods: []string{"Ol Kalou Town"}},
			{Name: "Kinangop", Wards: []string{"Engineer", "Gathabai", "North Kinangop", "Njabini/Kiburu"}, Neighborhoods: []string{"Njabini", "Engineer Town"}},
			{Name: "Ol Joro Orok", Wards: []string{"Gathanji", "Gatimu", "Weru"}, Neighborhoods: []string{"Ol Joro Orok"}},
			{Name: "Ndaragwa", Wards: []string{"Leshau Pondo", "Kiriita", "Central"}, Neighborhoods: []string{"Mairo Inya"}},
		},
	},
	{
		Code: 19, Name: "Nyeri",
		Towns: []string{"Nyeri", "Karatina", "Othaya", "Mukurweini", "Naro Moru"},
		SubCounties: []SubCountySeed{
			{Name: "Nyeri Town", Wards: []string{"Kiganjo/Mathari", "Rware", "Gatitu/Aguthi", "Ruring'u"}, Neighborhoods: []string{"Ruring'u", "Nyeri Town", "Ring Road"}},
			{Name: "Mathira East", Wards: []string{"Karatina Town", "Magutu", "Iria-ini"}, Neighborhoods: []string{"Karatina Town"}},
			{Name: "Othaya", Wards: []string{"Iria-ini", "Chinga", "Mahiga", "Karima"}, Neighborhoods: []string{"Othaya Town"}},
			{Name: "Mukurweini", Wards: []string{"Ruguru", "Gikondi", "Mukurweini West", "Mukurweini Central"}, Neighborhoods: []string{"Mukurweini Town"}},
			{Name: "Kieni East", Wards: []string{"Gakawa", "Naro Moru/Kiamathaga", "Thegu River"}, Neighborhoods: []string{"Naro Moru Town"}},
		},
	},
	{
		Code: 20, Name: "Kirinyaga",
		Towns: []string{"Kerugoya", "Sagana", "Kutus", "Wang'uru"},
		SubCounties: []SubCountySeed{
			{Name: "Kirinyaga Central", Wards: []string{"Mutira", "Kaelo", "Kerugoya"}, Neighborhoods: []string{"Kerugoya Town"}},
			{Name: "Mwea East", Wards: []string{"Tebere", "Nyangati", "Murinduko"}, Neighborhoods: []string{"Wang'uru Town"}},
			{Name: "Ndia", Wards: []string{"Kariti", "Mukure", "Sagana"}, Neighborhoods: []string{"Sagana Town"}},
			{Name: "Gichugu", Wards: []string{"Kabare", "Baragwi", "Njukiini"}, Neighborhoods: []string{"Kutus Town"}},
		},
	},
	{
		Code: 21, Name: "Murang'a",
		Towns: []string{"Murang'a", "Kenol", "Maragua", "Kiria-ini", "Kangema"},
		SubCounties: []SubCountySeed{
			{Name: "Murang'a South", Wards: []string{"Kimorori/Wempa", "Makuyu", "Kambiti"}, Neighborhoods: []string{"Kenol Town", "Makuyu"}},
			{Name: "Maragua", Wards: []string{"Ichagaki", "Nginda", "Makuyu"}, Neighborhoods: []string{"Maragua Town"}},
			{Name: "Kiharu", Wards: []string{"Township", "Mbiri", "Mugoiri"}, Neighborhoods: []string{"Murang'a Town"}},
			{Name: "Kangema", Wards: []string{"Kaniange", "Muguru", "Rwamuthambi"}, Neighborhoods: []string{"Kangema Town"}},
		},
	},
	{
		Code: 22, Name: "Kiambu",
		Towns: []string{"Thika", "Ruiru", "Kiambu", "Kikuyu", "Limuru", "Juja", "Karuri", "Githunguri"},
		SubCounties: []SubCountySeed{
			{Name: "Thika Town", Wards: []string{"Township", "Kamenu", "Hospital", "Gatuanyaga"}, Neighborhoods: []string{"Makongeni", "Section 9", "Landless", "Happy Valley"}},
			{Name: "Ruiru", Wards: []string{"Gitothua", "Biashara", "Gatongora", "Kahawa Sukari", "Kahawa Wendani", "Kiuu"}, Neighborhoods: []string{"Kahawa Sukari", "Kahawa Wendani", "Mwihoko", "Membley"}},
			{Name: "Juja", Wards: []string{"Juja", "Kalimoni", "Witeithie", "Theta"}, Neighborhoods: []string{"Juja Town", "HighPoint", "Gachororo"}},
			{Name: "Kiambu Town", Wards: []string{"Township", "Ndumberi", "Riabai", "Ting'ang'a"}, Neighborhoods: []string{"Kiambu Town", "Ndumberi", "Indian Bazaar"}},
			{Name: "Kabete", Wards: []string{"Gitaru", "Muguga", "Nyadhuna", "Kabete", "Uthiru"}, Neighborhoods: []string{"Uthiru", "King'eero", "Lower Kabete"}},
			{Name: "Kikuyu", Wards: []string{"Karai", "Kikuyu", "Sigona", "Kinoo"}, Neighborhoods: []string{"Kinoo", "Muthiga", "Kikuyu Town"}},
			{Name: "Limuru", Wards: []string{"Limuru Central", "Ndeiya", "Limuru East", "Ngecha"}, Neighborhoods: []string{"Limuru Town", "Ngecha"}},
			{Name: "Githunguri", Wards: []string{"Githunguri", "Githiga", "Ikinu", "Ngewa"}, Neighborhoods: []string{"Githunguri Town", "Ikinu"}},
		},
	},
	{
		Code: 23, Name: "Turkana",
		Towns: []string{"Lodwar", "Kakuma", "Lokichogio", "Lokichar"},
		SubCounties: []SubCountySeed{
			{Name: "Turkana Central", Wards: []string{"Kerio Delta", "Kang'atotha", "Lodwar Township"}, Neighborhoods: []string{"Lodwar Town"}},
			{Name: "Turkana West", Wards: []string{"Kakuma", "Lopur", "Letea"}, Neighborhoods: []string{"Kakuma Town"}},
			{Name: "Turkana South", Wards: []string{"Lokichar", "Katilu", "Kaputir"}, Neighborhoods: []string{"Lokichar Town"}},
		},
	},
	{
		Code: 24, Name: "West Pokot",
		Towns: []string{"Kapenguria", "Chepareria", "Makutano"},
		SubCounties: []SubCountySeed{
			{Name: "Kapenguria", Wards: []string{"Riwo", "Kapenguria", "Endugh"}, Neighborhoods: []string{"Kapenguria Town", "Makutano"}},
			{Name: "Pokot South", Wards: []string{"Lelan", "Chepareria", "Batei"}, Neighborhoods: []string{"Chepareria Town"}},
		},
	},
	{
		Code: 25, Name: "Samburu",
		Towns: []string{"Maralal", "Baragoi", "Wamba", "Archers Post"},
		SubCounties: []SubCountySeed{
			{Name: "Samburu West", Wards: []string{"Maralal", "Loosuk", "Poro"}, Neighborhoods: []string{"Maralal Town"}},
			{Name: "Samburu East", Wards: []string{"Wamba West", "Wamba East", "Waso"}, Neighborhoods: []string{"Wamba Town", "Archers Post"}},
		},
	},
	{
		Code: 26, Name: "Trans Nzoia",
		Towns: []string{"Kitale", "Kiminini", "Endebess"},
		SubCounties: []SubCountySeed{
			{Name: "Saboti", Wards: []string{"Kitale Township", "Matisi", "Tuwani", "Saboti"}, Neighborhoods: []string{"Matisi", "Tuwani", "Milimani Kitale"}},
			{Name: "Kiminini", Wards: []string{"Kiminini", "Waitaluk", "Sirende"}, Neighborhoods: []string{"Kiminini Town"}},
			{Name: "Endebess", Wards: []string{"Endebess", "Matumbei", "Chepchoina"}, Neighborhoods: []string{"Endebess Town"}},
		},
	},
	{
		Code: 27, Name: "Uasin Gishu",
		Towns: []string{"Eldoret", "Turbo", "Burnt Forest", "Moiben"},
		SubCounties: []SubCountySeed{
			{Name: "Ainabkoi", Wards: []string{"Kapsoya", "Kaptagat", "Ainabkoi"}, Neighborhoods: []string{"Kapsoya Estate", "Elgon View", "Kimumu"}},
			{Name: "Kapseret", Wards: []string{"Langas", "Simat/Kapseret", "Ngeria", "Megun"}, Neighborhoods: []string{"Langas Estate", "Pioneer", "Pioneer Estate"}},
			{Name: "Turbo", Wards: []string{"Ngenyilel", "Tapsagoi", "Kamagut", "Huruma"}, Neighborhoods: []string{"Huruma Estate", "Turbo Town"}},
			{Name: "Moiben", Wards: []string{"Tembelio", "Sergoit", "Karuna/Meibeki"}, Neighborhoods: []string{"Moiben Town"}},
		},
	},
	{
		Code: 28, Name: "Elgeyo/Marakwet",
		Towns: []string{"Iten", "Kapsowar", "Chebiemit"},
		SubCounties: []SubCountySeed{
			{Name: "Keiyo North", Wards: []string{"Emsoo", "Kamariny", "Tambach"}, Neighborhoods: []string{"Iten Town", "Tambach"}},
			{Name: "Marakwet West", Wards: []string{"Kapsowar", "Lelan", "Sengwer"}, Neighborhoods: []string{"Kapsowar Town"}},
		},
	},
	{
		Code: 29, Name: "Nandi",
		Towns: []string{"Kapsabet", "Nandi Hills", "Mosoriot"},
		SubCounties: []SubCountySeed{
			{Name: "Emgwen", Wards: []string{"Kapsabet", "Chepterwai", "Kilibwoni"}, Neighborhoods: []string{"Kapsabet Town"}},
			{Name: "Nandi Hills", Wards: []string{"Nandi Hills", "Chepkunyuk", "Ol'lessos"}, Neighborhoods: []string{"Nandi Hills Town"}},
			{Name: "Chesumei", Wards: []string{"Mosoriot", "Chemundu", "Kosirai"}, Neighborhoods: []string{"Mosoriot Town"}},
		},
	},
	{
		Code: 30, Name: "Baringo",
		Towns: []string{"Kabarnet", "Eldama Ravine", "Marigat"},
		SubCounties: []SubCountySeed{
			{Name: "Baringo Central", Wards: []string{"Kabarnet", "Sacho", "Tenges"}, Neighborhoods: []string{"Kabarnet Town"}},
			{Name: "Eldama Ravine", Wards: []string{"Lembus", "Ravine", "Mumberes"}, Neighborhoods: []string{"Eldama Ravine Town"}},
			{Name: "Baringo South", Wards: []string{"Marigat", "Ilchamus", "Mochongoi"}, Neighborhoods: []string{"Marigat Town"}},
		},
	},
	{
		Code: 31, Name: "Laikipia",
		Towns: []string{"Nanyuki", "Nyahururu", "Rumuruti"},
		SubCounties: []SubCountySeed{
			{Name: "Laikipia East", Wards: []string{"Nanyuki", "Umande", "Thingithu"}, Neighborhoods: []string{"Nanyuki Town", "Thingithu Estate"}},
			{Name: "Laikipia West", Wards: []string{"Nyahururu", "Rumuruti Township", "Githiga"}, Neighborhoods: []string{"Nyahururu Town", "Rumuruti"}},
		},
	},
	{
		Code: 32, Name: "Nakuru",
		Towns: []string{"Nakuru", "Naivasha", "Gilgil", "Molo", "Njoro"},
		SubCounties: []SubCountySeed{
			{Name: "Nakuru Town East", Wards: []string{"Biashara", "Kivumbini", "Flamingo", "Menengai", "Nakuru East"}, Neighborhoods: []string{"Section 58", "Milimani Nakuru", "Free Area", "Lanet"}},
			{Name: "Nakuru Town West", Wards: []string{"Barut", "London", "Kaptembwo", "Kipkenyo", "Rhonda"}, Neighborhoods: []string{"Kaptembwo", "Rhonda", "London Estate"}},
			{Name: "Naivasha", Wards: []string{"Naivasha East", "Viwandani", "Mai Mahiu", "Maeilla", "Olkaria"}, Neighborhoods: []string{"Naivasha Town", "Mai Mahiu", "Karagita"}},
			{Name: "Gilgil", Wards: []string{"Gilgil", "Elementaita", "Mbaruk/Eburu"}, Neighborhoods: []string{"Gilgil Town"}},
			{Name: "Molo", Wards: []string{"Molo", "Mariashoni", "Elburgon"}, Neighborhoods: []string{"Molo Town", "Elburgon"}},
		},
	},
	{
		Code: 33, Name: "Narok",
		Towns: []string{"Narok", "Kilgoris"},
		SubCounties: []SubCountySeed{
			{Name: "Narok North", Wards: []string{"Narok Town", "Nkareta", "Olokurto"}, Neighborhoods: []string{"Narok Town"}},
			{Name: "Trans Mara West", Wards: []string{"Kilgoris Central", "Keyian", "Angata Barikoi"}, Neighborhoods: []string{"Kilgoris Town"}},
		},
	},
	{
		Code: 34, Name: "Kajiado",
		Towns: []string{"Ngong", "Kitengela", "Ongata Rongai", "Kajiado", "Kiserian"},
		SubCounties: []SubCountySeed{
			{Name: "Kajiado North", Wards: []string{"Ngong", "Ongata Rongai", "Nkaimurunya", "Oloolua"}, Neighborhoods: []string{"Ongata Rongai", "Ngong Town", "Kiserian", "Matasia"}},
			{Name: "Kajiado East", Wards: []string{"Kitengela", "Oloosirkon/Sholinke", "Kenyewa-Poka"}, Neighborhoods: []string{"Kitengela Town", "Acacia Estate"}},
			{Name: "Kajiado Central", Wards: []string{"Purko", "Ildamat", "Dalalekutuk"}, Neighborhoods: []string{"Kajiado Town"}},
		},
	},
	{
		Code: 35, Name: "Kericho",
		Towns: []string{"Kericho", "Kipkelion", "Litein"},
		SubCounties: []SubCountySeed{
			{Name: "Ainamoi", Wards: []string{"Ainamoi", "Kapsoit", "Kipchebor"}, Neighborhoods: []string{"Kericho Town", "Kapsoit"}},
			{Name: "Bureti", Wards: []string{"Cheplanget", "Litein", "Cheboin"}, Neighborhoods: []string{"Litein Town"}},
		},
	},
	{
		Code: 36, Name: "Bomet",
		Towns: []string{"Bomet", "Sotik"},
		SubCounties: []SubCountySeed{
			{Name: "Bomet Central", Wards: []string{"Siloam", "Ndaraweta", "Singorwet"}, Neighborhoods: []string{"Bomet Town"}},
			{Name: "Sotik", Wards: []string{"Ndanai/Abosi", "Chemagel", "Manaret/Rongena"}, Neighborhoods: []string{"Sotik Town"}},
		},
	},
	{
		Code: 37, Name: "Kakamega",
		Towns: []string{"Kakamega", "Mumias", "Butere"},
		SubCounties: []SubCountySeed{
			{Name: "Lurambi", Wards: []string{"Butsotso East", "Butsotso South", "Sheywe"}, Neighborhoods: []string{"Kakamega Town", "Milimani Kakamega", "Amalemba"}},
			{Name: "Mumias West", Wards: []string{"Mumias Central", "Mumias North", "Etenje"}, Neighborhoods: []string{"Mumias Town"}},
			{Name: "Butere", Wards: []string{"Marama West", "Marama Central", "Marenyo"}, Neighborhoods: []string{"Butere Town"}},
		},
	},
	{
		Code: 38, Name: "Vihiga",
		Towns: []string{"Mbale", "Luanda", "Hamisi"},
		SubCounties: []SubCountySeed{
			{Name: "Vihiga", Wards: []string{"Lugaga-Wamuluma", "South Maragoli"}, Neighborhoods: []string{"Mbale Town"}},
			{Name: "Luanda", Wards: []string{"Luanda Township", "Wemilabi", "Emabungo"}, Neighborhoods: []string{"Luanda Town"}},
		},
	},
	{
		Code: 39, Name: "Bungoma",
		Towns: []string{"Bungoma", "Webuye", "Kimilili"},
		SubCounties: []SubCountySeed{
			{Name: "Kanduyi", Wards: []string{"Bungoma Township", "Bukembe West", "Khalaba"}, Neighborhoods: []string{"Bungoma Town", "Khalaba"}},
			{Name: "Webuye West", Wards: []string{"Sitikho", "Matulo", "Bokoli"}, Neighborhoods: []string{"Webuye Town"}},
			{Name: "Kimilili", Wards: []string{"Kimilili", "Kibingei", "Maeni"}, Neighborhoods: []string{"Kimilili Town"}},
		},
	},
	{
		Code: 40, Name: "Busia",
		Towns: []string{"Busia", "Malaba", "Nambale"},
		SubCounties: []SubCountySeed{
			{Name: "Matayos", Wards: []string{"Busia Township", "Burumba", "Mayanje"}, Neighborhoods: []string{"Busia Town", "Burumba Estate"}},
			{Name: "Teso North", Wards: []string{"Malaba Central", "Malaba North", "Ang'urai"}, Neighborhoods: []string{"Malaba Town"}},
		},
	},
	{
		Code: 41, Name: "Siaya",
		Towns: []string{"Siaya", "Bondo", "Ugunja"},
		SubCounties: []SubCountySeed{
			{Name: "Alego Usonga", Wards: []string{"Siaya Township", "Usonga", "North Alego"}, Neighborhoods: []string{"Siaya Town"}},
			{Name: "Bondo", Wards: []string{"Yimbo West", "Central Sakwa", "Bondo Township"}, Neighborhoods: []string{"Bondo Town"}},
			{Name: "Ugunja", Wards: []string{"Ugunja", "Sidindi", "Sigomere"}, Neighborhoods: []string{"Ugunja Town"}},
		},
	},
	{
		Code: 42, Name: "Kisumu",
		Towns: []string{"Kisumu City", "Muhoroni", "Ahero"},
		SubCounties: []SubCountySeed{
			{Name: "Kisumu Central", Wards: []string{"Railways", "Migosi", "Shaurimoyo", "Kondele", "Market Milimani"}, Neighborhoods: []string{"Milimani Kisumu", "Migosi", "Kondele", "Riat Hills", "Mamboleo"}},
			{Name: "Kisumu East", Wards: []string{"Kajulu", "Kolwa East", "Manyatta B"}, Neighborhoods: []string{"Manyatta Estate", "Buoye"}},
			{Name: "Nyando", Wards: []string{"Ahero", "East Kano", "Awasi/Onjiko"}, Neighborhoods: []string{"Ahero Town"}},
		},
	},
	{
		Code: 43, Name: "Homa Bay",
		Towns: []string{"Homa Bay", "Mbita", "Oyugis"},
		SubCounties: []SubCountySeed{
			{Name: "Homa Bay Town", Wards: []string{"Homa Bay Central", "Homa Bay Arujo", "Homa Bay West"}, Neighborhoods: []string{"Homa Bay Town", "Arujo Estate"}},
			{Name: "Suba North", Wards: []string{"Mbita", "Rusinga Island", "Gembe"}, Neighborhoods: []string{"Mbita Town"}},
			{Name: "Kasipul", Wards: []string{"Oyugis", "West Kasipul", "South Kasipul"}, Neighborhoods: []string{"Oyugis Town"}},
		},
	},
	{
		Code: 44, Name: "Migori",
		Towns: []string{"Migori", "Rongo", "Awendo", "Kehancha"},
		SubCounties: []SubCountySeed{
			{Name: "Suna West", Wards: []string{"Wiga", "Wasweta II", "Ragana-Oruba"}, Neighborhoods: []string{"Migori Town", "Oruba Estate"}},
			{Name: "Rongo", Wards: []string{"Central Kamagambo", "East Kamagambo"}, Neighborhoods: []string{"Rongo Town"}},
			{Name: "Awendo", Wards: []string{"North Sakwa", "Central Sakwa"}, Neighborhoods: []string{"Awendo Town"}},
		},
	},
	{
		Code: 45, Name: "Kisii",
		Towns: []string{"Kisii", "Ogembo", "Suneka"},
		SubCounties: []SubCountySeed{
			{Name: "Nyaribari Chache", Wards: []string{"Bobaracho", "Kisii Central", "Keumbu"}, Neighborhoods: []string{"Kisii Town", "Milimani Kisii", "Jogoo Estate"}},
			{Name: "Bomachoge Chache", Wards: []string{"Ogembo Town", "Majoge"}, Neighborhoods: []string{"Ogembo Town"}},
		},
	},
	{
		Code: 46, Name: "Nyamira",
		Towns: []string{"Nyamira", "Nyansiongo"},
		SubCounties: []SubCountySeed{
			{Name: "West Mugirango", Wards: []string{"Nyamira Township", "Bogichora", "Bosamaro"}, Neighborhoods: []string{"Nyamira Town"}},
			{Name: "Borabu", Wards: []string{"Mekenene", "Nyansiongo", "Esise"}, Neighborhoods: []string{"Nyansiongo"}},
		},
	},
	{
		Code: 47, Name: "Nairobi City",
		Towns: []string{"Nairobi City"},
		SubCounties: []SubCountySeed{
			{Name: "Westlands", Wards: []string{"Kitisuru", "Parklands/Highridge", "Karura", "Kangemi", "Mountain View"}, Neighborhoods: []string{"Westlands", "Parklands", "Spring Valley", "Brookside", "Kangemi"}},
			{Name: "Dagoretti North", Wards: []string{"Kilimani", "Kawangware", "Gatina", "Kileleshwa", "Kabiro"}, Neighborhoods: []string{"Kilimani", "Kileleshwa", "Lavington", "Kawangware"}},
			{Name: "Dagoretti South", Wards: []string{"Mutu-ini", "Ngando", "Riruta", "Uthiru/Ruthimitu", "Waithaka"}, Neighborhoods: []string{"Riruta", "Waithaka", "Dagoretti Corner"}},
			{Name: "Lang'ata", Wards: []string{"Karen", "Nairobi West", "Mugumo-ini", "South C", "Nyayo Highrise"}, Neighborhoods: []string{"Karen", "South C", "Nairobi West", "Dam Estate", "Langata"}},
			{Name: "Kibra", Wards: []string{"Laini Saba", "Lindi", "Makina", "Woodley/Kenyatta Golf Course", "Sarang'ombe"}, Neighborhoods: []string{"Woodley", "Golf Course Estate", "Adams Arcade"}},
			{Name: "Roysambu", Wards: []string{"Githurai", "Kahawa West", "Zimmerman", "Roysambu", "Kahawa"}, Neighborhoods: []string{"Roysambu", "Kahawa West", "Zimmerman", "TRM Area"}},
			{Name: "Kasarani", Wards: []string{"Clay City", "Mwiki", "Kasarani", "Njiru", "Ruai"}, Neighborhoods: []string{"Kasarani", "Mwiki", "Clay City", "Ruai", "Njiru"}},
			{Name: "Ruaraka", Wards: []string{"Babadogo", "Utalii", "Mathare North", "Lucky Summer", "Korogocho"}, Neighborhoods: []string{"Baba Dogo", "Lucky Summer", "Garden Estate"}},
			{Name: "Embakasi South", Wards: []string{"Imara Daima", "Kwa Njenga", "Kwa Reuben", "Pipeline", "Kware"}, Neighborhoods: []string{"Imara Daima", "Pipeline Estate"}},
			{Name: "Embakasi North", Wards: []string{"Kariobangi North", "Dandora Area I", "Dandora Area II", "Dandora Area III", "Dandora Area IV"}, Neighborhoods: []string{"Dandora", "Kariobangi"}},
			{Name: "Embakasi Central", Wards: []string{"Kayole North", "Kayole Central", "Kayole South", "Komarock", "Matopeni/Spring Valley"}, Neighborhoods: []string{"Komarock", "Kayole"}},
			{Name: "Embakasi East", Wards: []string{"Upper Savanna", "Lower Savanna", "Embakasi", "Utawala", "Mihango"}, Neighborhoods: []string{"Utawala", "Embakasi Village", "Mihango"}},
			{Name: "Embakasi West", Wards: []string{"Umoja I", "Umoja II", "Mowlem", "Donholm"}, Neighborhoods: []string{"Donholm", "Umoja", "Tena Estate", "Fedha"}},
			{Name: "Makadara", Wards: []string{"Maringo/Hamza", "Harambee", "Makadara", "Viwandani"}, Neighborhoods: []string{"Buruburu", "Makadara", "Hamza"}},
			{Name: "Kamukunji", Wards: []string{"Pumwani", "Eastleigh North", "Eastleigh South", "Airbase", "California"}, Neighborhoods: []string{"Eastleigh", "Pumwani"}},
			{Name: "Starehe", Wards: []string{"Nairobi Central", "Ngara", "Pangani", "Ziwani/Kariokor", "Landimawe", "Nairobi South"}, Neighborhoods: []string{"Nairobi CBD", "Ngara", "South B", "Pangani", "Park Road"}},
			{Name: "Mathare", Wards: []string{"Hospital", "Mabatini", "Huruma", "Ngei", "Mlango Kubwa", "Kiamaiko"}, Neighborhoods: []string{"Huruma", "Mathare"}},
		},
	},
}
