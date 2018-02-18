package main

import rl "github.com/chzyer/readline"

/* defaultGetCompletion
is a parameter list of the default parameters for the get command. This
list will be added to every completer */
var defaultGetCompletion = rl.PcItem("get",
	rl.PcItem("filter"),
	rl.PcItem("hosts"),
	rl.PcItem("selhosts"),
	rl.PcItem("allhosts"))

/* defaultSetCompletion
is a parameter list of the default parameters for the get command. This
list will be added to every completer */
var defaultSetCompletion = rl.PcItem("set",
	rl.PcItem("filter"))

/* cliNetironCompleter
is an autocompletion tree for the Netiron command line
*/
var cliNetironCompleter = rl.NewPrefixCompleter(
	rl.PcItem("show",
		rl.PcItem("access-list"),
		rl.PcItem("acl-policy"),
		rl.PcItem("arp"),
		rl.PcItem("bfd",
			rl.PcItem("applications"),
			rl.PcItem("mpls"),
			rl.PcItem("neighbors"),
			rl.PcItem("neighbors bgp"),
			rl.PcItem("neighbors details"),
			rl.PcItem("neighbors interface"),
			rl.PcItem("neighbors isis"),
			rl.PcItem("neighbors ospf"),
			rl.PcItem("neighbors ospf"),
			rl.PcItem("neighbors static"),
			rl.PcItem("neighbors static")),
		rl.PcItem("chassis"),
		rl.PcItem("configuration "),
		rl.PcItem("cpu histogram"),
		rl.PcItem("interface ethernet"),
		rl.PcItem("interfaces tunnel"),
		rl.PcItem("ip",
			rl.PcItem("bgp", rl.PcItem("attribute-entries"),
				rl.PcItem("config"),
				rl.PcItem("dampened-paths"),
				rl.PcItem("filtered-routes "),
				rl.PcItem("flap-statistics "),
				rl.PcItem("ipv6"),
				rl.PcItem("neighbors "),
				rl.PcItem("neighbors advertised-routes"),
				rl.PcItem("neighbors flap-statistics "),
				rl.PcItem("neighbors last-packet-with-error "),
				rl.PcItem("neighbors received "),
				rl.PcItem("neighbors received-routes"),
				rl.PcItem("neighbors rib-out-routes"),
				rl.PcItem("routes community "),
				rl.PcItem("neighbors routes "),
				rl.PcItem("neighbors routes-summary"),
				rl.PcItem("peer-group"),
				rl.PcItem("routes"),
				rl.PcItem("summary"),
				rl.PcItem("vrf neighbors "),
				rl.PcItem("vrf routes "),
				rl.PcItem("vrf ")),
			rl.PcItem("interface"),
			rl.PcItem("mbgp ipv6"),
			rl.PcItem("multicast"),
			rl.PcItem("multicast vpls"),
			rl.PcItem("ospf"),
			rl.PcItem("route"),
			rl.PcItem("static-arp"),
			rl.PcItem("vrrp"),
			rl.PcItem("vrrp-extended")),
		rl.PcItem("ipsec",
			rl.PcItem("egress-config"),
			rl.PcItem("egress-spi-table"),
			rl.PcItem("error-count"),
			rl.PcItem("ingress-config"),
			rl.PcItem("ingress-spi-table"),
			rl.PcItem("policy"),
			rl.PcItem("profile"),
			rl.PcItem("proposal"),
			rl.PcItem("sa"),
			rl.PcItem("statistics")),
		rl.PcItem("ip-tunnels"),
		rl.PcItem("ipv6",
			rl.PcItem("access-list bindings "),
			rl.PcItem("access-list receive accounting "),
			rl.PcItem("bgp"),
			rl.PcItem("bgp neighbors"),
			rl.PcItem("bgp routes "),
			rl.PcItem("bgp summary"),
			rl.PcItem("dhcp-relay interface"),
			rl.PcItem("dhcp-relay options"),
			rl.PcItem("interface tunnel"),
			rl.PcItem("ospf interface "),
			rl.PcItem("vrrp"),
			rl.PcItem("vrrp-extended")),
		rl.PcItem("isis"),
		rl.PcItem("license"),
		rl.PcItem("log"),
		rl.PcItem("module"),
		rl.PcItem("mpls",
			rl.PcItem("autobw-threshold-table "),
			rl.PcItem("bypass-lsp"),
			rl.PcItem("config"),
			rl.PcItem("forwarding"),
			rl.PcItem("interface"),
			rl.PcItem("label-range"),
			rl.PcItem("ldp"),
			rl.PcItem("ldp database"),
			rl.PcItem("ldp fec"),
			rl.PcItem("ldp interface"),
			rl.PcItem("ldp neighbor"),
			rl.PcItem("ldp path"),
			rl.PcItem("ldp peer"),
			rl.PcItem("ldp session "),
			rl.PcItem("ldp statistics"),
			rl.PcItem("ldp tunnel "),
			rl.PcItem("lsp"),
			rl.PcItem("lsp_pmp_xc "),
			rl.PcItem("path"),
			rl.PcItem("policy "),
			rl.PcItem("route "),
			rl.PcItem("rsvp",
				rl.PcItem("interface"),
				rl.PcItem("neighbor"),
				rl.PcItem("session"),
				rl.PcItem("session backup"),
				rl.PcItem("session brief"),
				rl.PcItem("session bypass"),
				rl.PcItem("session destination"),
				rl.PcItem("session detail"),
				rl.PcItem("session detour"),
				rl.PcItem("session down"),
				rl.PcItem("session extensive"),
				rl.PcItem("session (ingress/egress)"),
				rl.PcItem("session (interface)"),
				rl.PcItem("session name"),
				rl.PcItem("session pmp"),
				rl.PcItem("session pp"),
				rl.PcItem("session ppend"),
				rl.PcItem("session transit"),
				rl.PcItem("session up"),
				rl.PcItem("session wide"),
				rl.PcItem("statistics")),
			rl.PcItem("static-lsp"),
			rl.PcItem("statistics",
				rl.PcItem("pe"),
				rl.PcItem("bypass-lsp"),
				rl.PcItem("label"),
				rl.PcItem("ldp transit"),
				rl.PcItem("ldp tunnel "),
				rl.PcItem("lsp"),
				rl.PcItem("oam"),
				rl.PcItem("vll"),
				rl.PcItem("vll-local"),
				rl.PcItem("vpls"),
				rl.PcItem("vrf")),
			rl.PcItem("summary"),
			rl.PcItem("ted database"),
			rl.PcItem("ted path"),
			rl.PcItem("vll"),
			rl.PcItem("vll-local"),
			rl.PcItem("vpls")),
		rl.PcItem("openflow",
			rl.PcItem("controller"),
			rl.PcItem("flows"),
			rl.PcItem("groups"),
			rl.PcItem("interface"),
			rl.PcItem("meters"),
			rl.PcItem("queues")),
		rl.PcItem("rate-limit",
			rl.PcItem("counters bum-drop"),
			rl.PcItem("detail"),
			rl.PcItem("interface"),
			rl.PcItem("ipv6 hoplimit-expired-to-cpu"),
			rl.PcItem("option-pkt-to-cpu"),
			rl.PcItem("ttl-expired-to-cpu")),
		rl.PcItem("route-map"),
		rl.PcItem("running-config"),
		rl.PcItem("sflow statistics "),
		rl.PcItem("spanning-tree "),
		rl.PcItem("statistics "),
		rl.PcItem("terminal"),
		rl.PcItem("version"),
		rl.PcItem("vlan"),
	),
	defaultGetCompletion,
	defaultSetCompletion,
)

/* cliJunOSCompleter
is an autocompletion tree for the Netiron command line
*/
var cliJunOSCompleter = rl.NewPrefixCompleter(
	rl.PcItem("show",
		rl.PcItem("arp", rl.PcItem("no-resolve")),
		rl.PcItem("bfd"),
		rl.PcItem("bgp",
			rl.PcItem("group"),
			rl.PcItem("neighbor"),
			rl.PcItem("summary"),
		),
		rl.PcItem("chassis",
			rl.PcItem("alarms"),
			rl.PcItem("environment"),
			rl.PcItem("firmware"),
			rl.PcItem("hardware"),
			rl.PcItem("location"),
			rl.PcItem("pic"),
			rl.PcItem("pic-mode"),
			rl.PcItem("routing-engine"),
		),
		rl.PcItem("ethernet-switching",
			rl.PcItem("filters"),
			rl.PcItem("interfaces"),
			rl.PcItem("next-hops"),
			rl.PcItem("statistics"),
			rl.PcItem("table"),
		),
		rl.PcItem("firewall",
			rl.PcItem("application"),
			rl.PcItem("counter"),
			rl.PcItem("filter"),
			rl.PcItem("log"),
			rl.PcItem("terse"),
		),
		rl.PcItem("igmp"),
		rl.PcItem("interfaces"),
		rl.PcItem("ipv6",
			rl.PcItem("neighbors"),
			rl.PcItem("router-advertisement"),
		),
		rl.PcItem("lacp",
			rl.PcItem("interfaces"),
			rl.PcItem("statistics"),
			rl.PcItem("timeouts"),
		),
		rl.PcItem("lldp",
			rl.PcItem("detail"),
			rl.PcItem("local-information"),
			rl.PcItem("neighbors"),
			rl.PcItem("statistics"),
		),
		rl.PcItem("show"),
		rl.PcItem("mpls",
			rl.PcItem("interface"),
			rl.PcItem("lsp"),
			rl.PcItem("path"),
		),
		rl.PcItem("ospf",
			rl.PcItem("database"),
			rl.PcItem("interface"),
			rl.PcItem("log"),
			rl.PcItem("neighbor",
				rl.PcItem("area"),
				rl.PcItem("brief"),
				rl.PcItem("detail"),
				rl.PcItem("extensive"),
				rl.PcItem("instance"),
				rl.PcItem("interface"),
			),
			rl.PcItem("overview"),
			rl.PcItem("route"),
			rl.PcItem("statistics"),
		),
		rl.PcItem("ospf3",
			rl.PcItem("database"),
			rl.PcItem("interface"),
			rl.PcItem("log"),
			rl.PcItem("neighbor",
				rl.PcItem("area"),
				rl.PcItem("brief"),
				rl.PcItem("detail"),
				rl.PcItem("extensive"),
				rl.PcItem("instance"),
				rl.PcItem("interface"),
			),
			rl.PcItem("overview"),
			rl.PcItem("route"),
			rl.PcItem("statistics"),
		),
		rl.PcItem("route",
			rl.PcItem("advertising-protocol"),
			rl.PcItem("best"),
			rl.PcItem("brief"),
			rl.PcItem("detail"),
			rl.PcItem("instance"),
			rl.PcItem("martians"),
			rl.PcItem("next-hop"),
			rl.PcItem("protocol",
				rl.PcItem("arp"),
				rl.PcItem("bgp"),
				rl.PcItem("direct"),
				rl.PcItem("isis"),
				rl.PcItem("local"),
				rl.PcItem("mpls"),
				rl.PcItem("ospf"),
				rl.PcItem("ospf2"),
				rl.PcItem("ospf3"),
				rl.PcItem("static"),
			),
		),
		rl.PcItem("sflow"),
		rl.PcItem("snmp"),
		rl.PcItem("spanning-tree"),
		rl.PcItem("system",
			rl.PcItem("alarms",
				rl.PcItem("arp"),
				rl.PcItem("commit"),
				rl.PcItem("configurations"),
				rl.PcItem("connections"),
				rl.PcItem("license"),
				rl.PcItem("login"),
				rl.PcItem("memory"),
				rl.PcItem("processes"),
				rl.PcItem("reboot"),
				rl.PcItem("services"),
				rl.PcItem("software"),
				rl.PcItem("storage"),
				rl.PcItem("uptime"),
				rl.PcItem("users"),
			)),
		rl.PcItem("version"),
		rl.PcItem("virtual-chassis"),
		rl.PcItem("vlans"),
		rl.PcItem("vrrp"),
	),
	defaultGetCompletion,
	defaultSetCompletion,
)
