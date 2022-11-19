-- organisations

INSERT INTO organisations (id, name) VALUES ('5599931a-6f8e-442d-b52e-8c297af7cb8e', 'Bob E.V.');
INSERT INTO organisations (id, name) VALUES ('e9b6d487-0fd8-49ee-a48a-8afc63362ed3', 'MacLarens');
INSERT INTO organisations (id, name) VALUES ('25d638a8-3238-46b5-80c8-489b396de8a4', 'Brown University');
INSERT INTO organisations (id, name) VALUES ('8f8e2b3b-7420-4883-adae-2877aebc3345', 'OECD Environment Working Papers');
INSERT INTO organisations (id, name) VALUES ('55999312-6f8e-442d-b52e-8c297af7cb8e', 'Indian Institute of Science');
INSERT INTO organisations (id, name) VALUES ('027bfef3-73e6-4f1f-8595-1fb8dfc8bbe6', 'University of Sri Jayewardenepura');
INSERT INTO organisations (id, name) VALUES ('f8262283-0a55-4ab1-a862-cdc7daf77449', 'Universidad Autónoma de Nayarit');
INSERT INTO organisations (id, name) VALUES ('633081a3-5755-4f24-b68f-94f55588392a', 'Université de Toulouse');

-- users

INSERT INTO users (id, name, organisation) VALUES ('4e805cc9-fe3b-4649-96fc-f39634a557cd', 'Bob', '5599931a-6f8e-442d-b52e-8c297af7cb8e');
INSERT INTO users (id, name) VALUES ('c85e9879-f808-450d-8ab3-c6f5ab0e9d0c', 'Sponge-Bob');

INSERT INTO users (name, organisation) VALUES ('Barney Stinson', 'e9b6d487-0fd8-49ee-a48a-8afc63362ed3');
INSERT INTO users (name, organisation) VALUES ('Ted Mosby', 'e9b6d487-0fd8-49ee-a48a-8afc63362ed3');
INSERT INTO users (name, organisation) VALUES ('Robin Scherbatsky', 'e9b6d487-0fd8-49ee-a48a-8afc63362ed3');
INSERT INTO users (name, organisation) VALUES ('Lily Aldrin', 'e9b6d487-0fd8-49ee-a48a-8afc63362ed3');
INSERT INTO users (name, organisation) VALUES ('Marshall Eriksen', 'e9b6d487-0fd8-49ee-a48a-8afc63362ed3');

INSERT INTO users (name, organisation) VALUES ('B.K. Rosen', '25d638a8-3238-46b5-80c8-489b396de8a4');
INSERT INTO users (name, organisation) VALUES ('M.N. Wegman', '25d638a8-3238-46b5-80c8-489b396de8a4');
INSERT INTO users (name, organisation) VALUES ('F.K. Zadeck', '25d638a8-3238-46b5-80c8-489b396de8a4');

INSERT INTO users (name, organisation) VALUES ('Réka Soós', '8f8e2b3b-7420-4883-adae-2877aebc3345');
INSERT INTO users (name, organisation) VALUES ('Andrew Whiteman', '8f8e2b3b-7420-4883-adae-2877aebc3345');
INSERT INTO users (name, organisation) VALUES ('Gabriela Gavgas', '8f8e2b3b-7420-4883-adae-2877aebc3345');

INSERT INTO users (name, organisation) VALUES ('B.V. Venkatarama Reddy', '5599931a-6f8e-442d-b52e-8c297af7cb8e');
INSERT INTO users (name, organisation) VALUES ('K.S. Jagadish', '5599931a-6f8e-442d-b52e-8c297af7cb8e');

INSERT INTO users (name, organisation) VALUES ('Branavan Arulmoly', '027bfef3-73e6-4f1f-8595-1fb8dfc8bbe6');
INSERT INTO users (name, organisation) VALUES ('Chaminda Konthesingha', '027bfef3-73e6-4f1f-8595-1fb8dfc8bbe6');
INSERT INTO users (name, organisation) VALUES ('Anura Nanayakkara', '027bfef3-73e6-4f1f-8595-1fb8dfc8bbe6');

INSERT INTO users (name, organisation) VALUES ('Ulises Mercado Burciaga', 'f8262283-0a55-4ab1-a862-cdc7daf77449');

INSERT INTO users (name, organisation) VALUES ('Marie Chan', '633081a3-5755-4f24-b68f-94f55588392a');
INSERT INTO users (name, organisation) VALUES ('Eric Campo', '633081a3-5755-4f24-b68f-94f55588392a');
INSERT INTO users (name, organisation) VALUES ('Daniel Estève', '633081a3-5755-4f24-b68f-94f55588392a');
INSERT INTO users (name, organisation) VALUES ('AJean-Yves Fourniols', '633081a3-5755-4f24-b68f-94f55588392a');

-- securities

INSERT INTO securities (id, name, creator, description, ttl_1, ttl_2, funding_goal, funding_remaining) VALUES (
    '3e8b7701-9d3e-407a-b78a-d8fa4d07bff5', 
    'LMU',
    '5599931a-6f8e-442d-b52e-8c297af7cb8e',
    'Unexzellent',
    86400 * 120,
    86400 * 180,
    1000000000,
    1000000000
);

INSERT INTO securities (name, creator, description, ttl_1, ttl_2, funding_goal, funding_remaining) VALUES (
    'Global value numbers and redundant computations',
    '25d638a8-3238-46b5-80c8-489b396de8a4',
    'Most previous redundancy elilmination algorithms have been of two kinds. The lexical algorithms deal with the entire program, but they can only detect redundancy among computations of lexicatlly identical expressions, where expressions are lexically identical if they apply exactly the same operator to exactly the same operands. The value numbering algorithms,, on the other hand, can recognize redundancy among ex:pressions that are lexically different but that are certain to compute the same value. This is accomplished by assigning special symbolic names called value numbers to expr,essions. If the value numbers of the operands of two expressions are identical, and if the operators applied by the expressions are identical, then the expressions receive the: same value number and are certain to have the same values. Sameness of value numbers permits more extensive optimization than lexical identity, but value numbering algor:ithms have usually been restricted in the past to basic blocks (sequences of computations with no branching) or extended basic blocks (sequences of computations with no joins). We propose a redundancy elimination algorithm that is global (in that it deals with the entire program), yet able to recognize redundancy among expressions that are lexitally different.',
    86400 * 120,
    86400 * 180,
    100000000,
    100000000
);

INSERT INTO securities (name, creator, description, ttl_1, ttl_2, funding_goal, funding_remaining) VALUES (
    'Reducing the cost of preventing ocean plastic pollution',
    '8f8e2b3b-7420-4883-adae-2877aebc3345',
    'This paper provides estimates of the cost of preventing land-based plastic leakage into the ocean, covering 38 OECD member countries and 10 selected major plastic waste emitters in Asia and Africa. The study estimates capital costs at EUR 54 billion in the Moderate Ambition scenario and EUR 74 billion in the High Ambition scenario. The annualised per-capita costs range between EUR 0.2 to 6.5 in the Moderate Ambition scenario and from EUR 0.8 to 6.5 in the High Ambition scenario. These cost estimates are much lower than UNEP and ISWA estimates of the cost of inaction of inadequate waste management, roughly USD 9 to 45 per capita. Differences in estimated costs are found to depend on countries’ waste policy stringency and waste management infrastructure. This paper contributes to OECD work in support of a sustainable ocean economy and the Global Plastics Outlook report.',
    86400 * 40,
    86400 * 90,
    1000000,
    1000000
);

INSERT INTO securities (name, creator, description, ttl_1, ttl_2, funding_goal, funding_remaining) VALUES (
    'Improving embodied energy of common and alternative building materials and technologies',
    '5599931a-6f8e-442d-b52e-8c297af7cb8e',
    'Considerable amount of energy is spent in the manufacturing processes and transportation of various building materials. Conservation of energy becomes important in the context of limiting of green house gases emission into the atmosphere and reducing costs of materials. The paper is focused around some issues pertaining to embodied energy in buildings particularly in the Indian context. Energy consumption in the production of basic building materials (such as cement, steel, etc.) and different types of materials used for construction has been discussed. Energy spent in transportation of various building materials is presented. A comparison of energy in different types of masonry has been made. Energy in different types of alternative roofing systems has been discussed and compared with the energy of conventional reinforced concrete (RC) slab roof. Total embodied energy of a multi-storeyed building, a load bearing brickwork building and a soil-cement block building using alternative building materials has been compared. It has been shown that total embodied energy of load bearing masonry buildings can be reduced by 50% when energy efficient/alternative building materials are used.',
    86400 * 30,
    86400 * 180,
    100000000,
    100000000
);

INSERT INTO securities (name, creator, description, ttl_1, ttl_2, funding_goal, funding_remaining) VALUES (
    'New processes of cement mortar produced with manufactured sand and offshore sand as alternatives for river sand',
    '027bfef3-73e6-4f1f-8595-1fb8dfc8bbe6',
    'This study investigates the fresh and hardened state properties of cement-sand mortar comprising manufactured sand and offshore sand as alternatives for a complete replacement of river sand. Two types of manufactured sand were selected based on different rock types such as Hornblende-Gneiss and Charnockite. Offshore sand was collected from an open stock pile after required period of washing. Mortars were manufactured with a binder of Portland Limestone Cement. Binder-to-aggregate ratios of 1:3, 1:4 and 1:6 were considered in this study and manufactured sand was replaced at 0%, 25%, 50% and 75% with offshore sand. To check the influence of sand alternatives and blending ratios, fresh and hardened state properties of alternative mortars were analyzed and compared with reference mortars which were made with river sand alone. Wet and dry bulk densities of mortars were increased with lower replacement levels with offshore sand. Most mortars with blended sand improved the workability while consistency and initial setting time of mortars were not significantly affected. Inflated bleeding of mortars was noticed with the alternatives and replacement levels. Workable life was decreased at small replacements. When manufactured sand in mortar content was 25% and 50%, the water retentivity was significantly improved than other replacements and control mixes. Mortars at lower replacements greatly advanced the flexural strength, compressive strength and capillary water absorption. Linear shrinkage and thermal expansion of mortars were also affected with the selected replacement levels. Based on the overall performance of mortars, blended sand at 25% replacement of manufactured sand with offshore sand was deduced as the feasible solution for completely replacing river sand and to produce economical mortars.',
    86400 * 30,
    86400 * 180,
    750000000,
    750000000
);

INSERT INTO securities (name, creator, description, ttl_1, ttl_2, funding_goal, funding_remaining) VALUES (
    'Improving smart homes — Current features and future perspectives',
    'f8262283-0a55-4ab1-a862-cdc7daf77449',
    'In an ageing world, maintaining good health and independence for as long as possible is essential. Instead of hospitalization or institutionalization, the elderly and disabled can be assisted in their own environment 24 h a day with numerous ‘smart’ devices. The concept of the smart home is a promising and cost-effective way of improving home care for the elderly and the disabled in a non-obtrusive way, allowing greater independence, maintaining good health and preventing social isolation. Smart homes are equipped with sensors, actuators, and/or biomedical monitors. The devices operate in a network connected to a remote centre for data collection and processing. The remote centre diagnoses the ongoing situation and initiates assistance procedures as required. The technology can be extended to wearable and in vivo implantable devices to monitor people 24 h a day both inside and outside the house. This review describes a selection of projects in developed countries on smart homes examining the various technologies available. Advantages and disadvantages, as well as the impact on modern society, are discussed. Finally, future perspectives on smart homes as part of a home-based health care network are presented.',
    86400 * 30,
    86400 * 180,
    5000000000,
    5000000000
);

INSERT INTO securities (name, creator, description, ttl_1, ttl_2, funding_goal, funding_remaining) VALUES (
    'Housing Building Organizations for the Design of Strategies against Climate Change',
    '633081a3-5755-4f24-b68f-94f55588392a',
    'One of the biggest problems facing humanity is climate change, and the construction industry is one of the sectors causing the greatest impact. Therefore, design strategies accompanied by new methodologies are necessary. In this sense, this paper aims to assess sustainability for the design of organizational strategies against climate change, based on a holistic and systemic approach to sustainability development, in order to contribute to the decision-making in housing building organizations. The assessment was based on: 1) climate change indicators were selected from a case study; 2) a survey based on climate change indicators was designed and applied to 21% of the total organizations under study; and 3) critical indicators were identified. The result shows that 58% of the climate change indicators are critical and give evidence of the negative outlook that housing building organizations have in terms of sustainability. About 69% of these indicators belong to the cultural dimension. This demonstrates the lack of knowledge, customs, habits, and commitment to implementing sustainable strategies against climate change in these organizations. Finally, the results can contribute to designing strategies to promote sustainable building by the local government, and thus achieve more sustainable organizations that contribute to reducing their impact on climate change.',
    86400 * 30,
    86400 * 180,
    50000000,
    50000000
);